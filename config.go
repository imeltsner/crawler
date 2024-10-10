package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(rawBaseURL string, maxConcurrencyControl int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing url: %v, %v", rawBaseURL, err)
		return nil, err
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrencyControl),
		wg:                 &sync.WaitGroup{},
	}, nil
}
