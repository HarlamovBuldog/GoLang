// Findlinks1 prints links to HTML-document
// read from default input
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	/*
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	*/
	for _, link2 := range visitAddTask(nil, doc) {
		fmt.Println(link2)
	}
}

// "visit" adds all links, that were found in Node "n",
// to "links" slice and returns result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func visitAddTask(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if c := n.FirstChild; c != nil {
		links = visitAddTask(visitAddTask(links, c), c.NextSibling)
	}
	return links
}
