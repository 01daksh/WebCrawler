package models


type CrawlerBO struct{
	SeedUrl string `json:"seedurl"`
}

func (bo *CrawlerBO) GetSeedUrl()string{
	return string(bo.SeedUrl)
}
