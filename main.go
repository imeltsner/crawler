package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("no website provided")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v", args[1])
}
