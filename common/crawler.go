package common

import (
	"sync"
)

type TraversalData struct {
	Queue   chan UrlInfo
	Visited sync.Map
	WaitGroup *sync.WaitGroup
	Links []string
}

type UrlInfo struct{
	SeedUrl string
	Level int
}

func NewTraversalData() *TraversalData {
	return &TraversalData{
		Queue: make(chan UrlInfo),
		WaitGroup: &sync.WaitGroup{},
		Links: make([]string, 0),
	}
}
