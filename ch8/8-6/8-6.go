package main

import (
	"fmt"
	"gopl-practise/ch5/5-13/links"
	"log"
	"os"
)

func main() {
	worklist := make(chan []work)
	var n int

	n++
	go func() {
		var works []work
		for _, url := range os.Args[1:] {
			works = append(works, work{url, 1})
		}
		worklist <- works
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		works := <-worklist
		for _, w := range works {
			if !seen[w.url] {
				seen[w.url] = true
				n++
				go func(w work) {
					worklist <- crawl(w)
				}(w)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

type work struct {
	url   string
	depth int
}

func crawl(w work) []work {
	fmt.Printf("depth: %d, url: %s\n", w.depth, w.url)

	if w.depth >= 3 {
		return nil
	}

	tokens <- struct{}{}
	urls, err := links.Extract(w.url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	var works []work
	for _, url := range urls {
		works = append(works, work{url, w.depth + 1})
	}
	return works
}
