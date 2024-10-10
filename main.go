package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		fmt.Printf("program expects 3 args, got %v", len(args)-1)
		return
	}

	maxConcurrencyControl, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("second arg must be int")
		return
	}

	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("third arg must be int")
		return
	}

	cfg, err := configure(args[1], maxConcurrencyControl, maxPages)
	if err != nil {
		fmt.Printf("config error %v", err)
		return
	}

	fmt.Printf("starting crawl of: %v\n", args[1])
	cfg.wg.Add(1)
	go cfg.crawlPage(args[1])
	cfg.wg.Wait()

	printReport(cfg.pages, args[1])
}
