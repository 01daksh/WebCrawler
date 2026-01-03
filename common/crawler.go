package common

import (
	"sync"

	"github.com/spf13/viper"
)

type TraversalData struct {
	Queue   chan UrlInfo
	Visited sync.Map
	WaitGroup *sync.WaitGroup
	LinkInfo []LinkInformation
	Mutex sync.Mutex
}

type LinkInformation struct{
	Link string
	Level int
}

type UrlInfo struct{
	SeedUrl string
	Level int
}

func NewTraversalData() *TraversalData {
	return &TraversalData{
		Queue: make(chan UrlInfo),
		WaitGroup: &sync.WaitGroup{},
		LinkInfo: make([]LinkInformation, 0),
	}
}


func GetMongoConnectionString(key string)string{
	return viper.GetString(key)
}