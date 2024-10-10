package main

import (
	"fmt"
	"sort"
)

type pageEntry struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	header := fmt.Sprintf(`
=============================
REPORT for %v
=============================
	`, baseURL)
	fmt.Println(header)

	sortedPages := sortPagesMap(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %v\n", page.count, page.url)
	}
}

func sortPagesMap(pages map[string]int) []pageEntry {
	var pageEntries []pageEntry
	for k, v := range pages {
		pageEntries = append(pageEntries, pageEntry{k, v})
	}

	sort.Slice(pageEntries, func(i, j int) bool {
		if pageEntries[i].count == pageEntries[j].count {
			return pageEntries[i].url < pageEntries[j].url
		}
		return pageEntries[i].count > pageEntries[j].count
	})

	return pageEntries
}
