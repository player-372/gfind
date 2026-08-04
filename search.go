package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func search(target string, rootdir string, ignoreCase bool, hidden bool, itemType string, maxDepth int, currentDepth int, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	searchTarget := target
	if ignoreCase {
		searchTarget = strings.ToLower(target)
	}

	//Read directory
	dir, err := os.ReadDir(rootdir)
	if err != nil {
		return
	}

	//Scan directory
	for _, item := range dir {

		itemName := item.Name()

		if !hidden && strings.HasPrefix(itemName, ".") {
			continue
		}

		isMatch := false
		if ignoreCase {
			isMatch = strings.Contains(strings.ToLower(itemName), searchTarget)
		} else {
			isMatch = strings.Contains(itemName, target)
		}

		if isMatch {
			switch {
			default:

			case itemType == "d" && item.IsDir():
				results <- filepath.Join(rootdir, itemName)

			case itemType == "f" && !item.IsDir():
				results <- filepath.Join(rootdir, itemName)

			case itemType == "":
				results <- filepath.Join(rootdir, itemName)

			}
		}

		//Rerun for directories
		if currentDepth != maxDepth {

			if item.IsDir() {
				wg.Add(1)
				go search(target, filepath.Join(rootdir, itemName), ignoreCase, hidden, itemType, maxDepth, currentDepth+1, wg, results)
			}

		}

	}

}
