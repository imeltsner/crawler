package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)

	nodes, err := html.Parse(htmlReader)
	if err != nil {
		fmt.Errorf("unable to parse html %w", err)
		return []string{}, nil
	}

	var links []string

	var parseLinkData = func(attr []html.Attribute) {
		for _, attr := range attr {
			if attr.Key == "href" {
				if !strings.Contains(attr.Val, "https://") {
					attr.Val = rawBaseURL + attr.Val
				}
				links = append(links, attr.Val)
				return
			}
		}
	}

	var getLinks func(node *html.Node)
	getLinks = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			parseLinkData(node.Attr)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			getLinks(child)
		}
	}

	getLinks(nodes)
	return links, nil
}
