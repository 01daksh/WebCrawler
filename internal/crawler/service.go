package crawler

import (
	"WebCrawler/common"
	"WebCrawler/internal/models"
	"context"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"WebCrawler/internal/interfaces"
)

type CrawlerService struct {
	repo interfaces.CrawlerRepoInterface
}

func NewCrawlerService(repo interfaces.CrawlerRepoInterface) *CrawlerService {
	return &CrawlerService{
		repo: repo,
	}
}

const (
	batchSize int = 100
)

func (cw *CrawlerService) AddCrawler(ctx context.Context, req models.CrawlerBO) ([]common.LinkInformation, error) {

	links := cw.TraversePages(req.GetSeedUrl())

	if len(links) == 0 {
		return nil, nil
	}

	for i := 0; i < len(links); i += batchSize {
		end := i + batchSize
		if end > len(links) {
			end = len(links)
		}

		batch := links[i:end]

		if err := cw.repo.AddCrawler(ctx, batch); err != nil {
			return nil, err
		}
	}

	return links, nil
}

func (cw *CrawlerService) TraversePages(seedUrl string) []common.LinkInformation {
	traversalData := common.NewTraversalData()

	for i := 0; i < common.WorkerSize; i++ {
		go worker(traversalData)
	}

	traversalData.WaitGroup.Add(1)
	traversalData.Queue <- common.UrlInfo{
		SeedUrl: seedUrl,
		Level:   0,
	}

	traversalData.WaitGroup.Wait()
	close(traversalData.Queue)

	return traversalData.LinkInfo
}

func worker(traversalData *common.TraversalData) {
	for url := range traversalData.Queue {
		func() {
			defer traversalData.WaitGroup.Done()

			crawlPage(url, traversalData)

			traversalData.Mutex.Lock()

			traversalData.LinkInfo = append(traversalData.LinkInfo, common.LinkInformation{
				Link:  url.SeedUrl,
				Level: url.Level,
			})

			traversalData.Mutex.Unlock()
		}()
	}
}

func crawlPage(url common.UrlInfo, traversalData *common.TraversalData) {
	if url.Level > common.MaxDepth {
		return
	}

	if _, ok := traversalData.Visited.LoadOrStore(url.SeedUrl, true); ok {
		return
	}

	links := fetchAndExtract(url.SeedUrl)

	for _, link := range links {
		if !isValidHTTPSLink(link) {
			continue
		}

		fmt.Println(link)
		traversalData.WaitGroup.Add(1)
		traversalData.Queue <- common.UrlInfo{
			Level:   url.Level + 1,
			SeedUrl: link,
		}
	}

}

func fetchAndExtract(url string) (links []string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to get page:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("ERROR: Failed to parse HTML:", err)
		return
	}

	links = visit(nil, doc)
	return links
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func isValidHTTPSLink(link string) bool {
	return strings.HasPrefix(link, "https://")
}
