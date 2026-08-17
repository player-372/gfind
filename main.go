package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
)

func main() {

	rootdir := flag.String("root-dir", ".", "Directory to start the search from")
	ignoreCase := flag.Bool("ignore-case", false, "Search case-insensitively (e.g. 'Main' matches 'main')")
	hidden := flag.Bool("hidden", false, "Include hidden files and directories")
	itemType := flag.String("type", "", "Filter by type: 'f' for files only, 'd' for directories only")
	maxDepth := flag.Int("max-depth", 0, "Maximum search depth (0 = unlimited)")
	threads := flag.Int("threads", runtime.NumCPU()/2, "Number of concurrent search workers")
	flag.Parse()
	target := flag.Args()

	if *itemType != "" && *itemType != "f" && *itemType != "d" {
		fmt.Println("Error: --type must be 'f' or 'd'")
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("Error: rerun with one argument")
	} else {

		var wg sync.WaitGroup
		results := make(chan string, 10000)
		sem := make(chan struct{}, *threads)

		wg.Add(1)
		go search(target[0], *rootdir, *ignoreCase, *hidden, *itemType, *maxDepth, 1, &wg, results, sem)

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
