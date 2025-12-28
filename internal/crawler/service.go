package crawler

import (
	"WebCrawler/common"
	"WebCrawler/internal/models"
	"fmt"
	"net/http"

	"golang.org/x/net/html"

	"github.com/gin-gonic/gin"
)

type CrawlerService struct {
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

func (cw *CrawlerService) AddCrawler(c *gin.Context) {

	var req models.CrawlerBO
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}

	links := cw.TraversePages(req.GetSeedUrl())

	if links != nil {
		c.JSON(http.StatusOK, links)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "nothing for now!",
	})

}

func (cw *CrawlerService) TraversePages(seedUrl string) []string {
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

	return traversalData.Links
}

func worker(traversalData *common.TraversalData) {
	for url := range traversalData.Queue {
		func() {
			defer traversalData.WaitGroup.Done()
			crawlPage(url, traversalData)
			traversalData.Links = append(traversalData.Links, url.SeedUrl)
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
		fmt.Print(link, " ")
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
