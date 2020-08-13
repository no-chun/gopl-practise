package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	nodes := ElementByTagName(doc, "head", "script")
	for _, n := range nodes {
		fmt.Println(n.Data)
		fmt.Println(n.Attr)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, data := range name {
			if doc.Data == data {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementByTagName(c, name...)...)
	}

	return nodes
}
