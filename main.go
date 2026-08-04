package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	rootdir := flag.String("rootdir", ".", "Root direcotry")
	ignoreCase := flag.Bool("ignore-case", false, "Ignore case")
	flag.Parse()
	target := flag.Args()

	if len(flag.Args()) != 1 {
		fmt.Println("Error: rerun with one argument")
	} else {

		results := make(chan string)
		var wg sync.WaitGroup

		wg.Add(1)
		go search(target[0], *rootdir, *ignoreCase, &wg, results)

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
