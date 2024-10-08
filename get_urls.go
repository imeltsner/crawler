package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base url: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	nodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("unable to parse html %v", err)
	}

	var urls []string
	var parseURLData = func(attr []html.Attribute) {
		for _, attr := range attr {
			if attr.Key == "href" {
				href, err := url.Parse(attr.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
					continue
				}
				resolvedURL := baseURL.ResolveReference(href)
				urls = append(urls, resolvedURL.String())
				return
			}
		}
	}

	var getLinks func(node *html.Node)
	getLinks = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			parseURLData(node.Attr)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			getLinks(child)
		}
	}

	getLinks(nodes)
	return urls, nil
}
