package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	rootdir := flag.String("rootdir", ".", "Root direcotry")
	ignoreCase := flag.Bool("ignore-case", false, "Ignore case")
	hidden := flag.Bool("hidden", false, "Search hidden items")
	itemType := flag.String("type", "", "Search for files (f) or directories (d)? Don't use flag for both")
	maxDepth := flag.Int("max-depth", 0, "Max depth, 0 for infinite")
	flag.Parse()
	target := flag.Args()

	if *itemType != "" && *itemType != "f" && *itemType != "d" {
		fmt.Println("Error: --type must be 'f' or 'd'")
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("Error: rerun with one argument")
	} else {

		results := make(chan string)
		var wg sync.WaitGroup

		wg.Add(1)
		go search(target[0], *rootdir, *ignoreCase, *hidden, *itemType, *maxDepth, 1, &wg, results)

		//Getting results
		go func() {
			wg.Wait()
			close(results)
		}()

		for item := range results {
			fmt.Println(item)
		}

	}
}
