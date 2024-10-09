package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Errorf("unable to parse base URL: %v", err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Errorf("unable to parse current URL: %v", err)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Errorf("unable to normalize current url %v", err)
		return
	}

	if _, seen := pages[normalizedURL]; seen {
		pages[normalizedURL]++
		return
	} else {
		pages[normalizedURL] = 1
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Errorf("error getting html from current url %v", err)
		return
	}
	fmt.Printf("Getting html from: %v\n", rawCurrentURL)

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Errorf("error getting urls from html: %v", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
