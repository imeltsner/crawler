package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("no website provided")
		return
	}

	if len(args) > 2 {
		fmt.Printf("too many arguments provided")
		return
	}

	var pages = map[string]int{}
	fmt.Printf("starting crawl of: %v\n", args[1])
	crawlPage(args[1], args[1], pages)
}
