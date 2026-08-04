package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func search(target string, rootdir string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	dir, err := os.ReadDir(rootdir)
	if err != nil {
		return
	}

	for _, item := range dir {

		if strings.Contains(item.Name(), target) {
			results <- filepath.Join(rootdir, item.Name())
		}

		if item.IsDir() {
			wg.Add(1)
			go search(target, filepath.Join(rootdir, item.Name()), wg, results)
		}

	}

}
