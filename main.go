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

	cfg, err := configure(args[1], 5)
	if err != nil {
		fmt.Printf("config error %v", err)
		return
	}

	fmt.Printf("starting crawl of: %v\n", args[1])
	cfg.wg.Add(1)
	go cfg.crawlPage(args[1])
	cfg.wg.Wait()
}
