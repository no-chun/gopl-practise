package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	words, images, err := CountWordsAndImages("https://news.qq.com/")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("images: %d\nwords: %d\n", images, words)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		restwords, restimages := countWordsAndImages(c)
		words += restwords
		images += restimages
	}
	return
}
