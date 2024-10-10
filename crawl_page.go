package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.reachedMax() {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("unable to parse current URL: %v", err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("unable to normalize current url %v", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html from current url: %v %v", rawCurrentURL, err)
		return
	}
	fmt.Printf("Getting html from: %v\n", rawCurrentURL)

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, seen := cfg.pages[normalizedURL]; seen {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) reachedMax() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) >= cfg.maxPages
}
