package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dir struct {
	id   int
	size int64
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	info := make(chan dir)
	var waitGroup sync.WaitGroup
	for id, root := range roots {
		waitGroup.Add(1)
		go walkDir(root, &waitGroup, id, info)
	}
	go func() {
		waitGroup.Wait()
		close(info)
	}()

	tick := time.Tick(500 * time.Millisecond)
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case dir, ok := <-info:
			if !ok {
				break loop
			}
			nfiles[dir.id]++
			nbytes[dir.id] += dir.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for id, root := range roots {
		fmt.Printf("%d files %.1f GB in %s\n",
			nfiles[id], float64(nbytes[id])/1e9, root)
	}
}

func walkDir(d string, n *sync.WaitGroup, root int, info chan<- dir) {
	defer n.Done()
	for _, entry := range dirents(d) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(d, entry.Name())
			go walkDir(subdir, n, root, info)
		} else {
			info <- dir{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
