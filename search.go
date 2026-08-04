package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func search(target string, rootdir string, ignoreCase bool, hidden bool, wg *sync.WaitGroup, results chan string) {
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
			results <- filepath.Join(rootdir, itemName)
		}

		//Rerun for directories
		if item.IsDir() {
			wg.Add(1)
			go search(target, filepath.Join(rootdir, itemName), ignoreCase, hidden, wg, results)
		}

	}

}
