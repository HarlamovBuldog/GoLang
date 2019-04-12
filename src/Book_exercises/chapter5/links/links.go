// Package links contains functions to work with links.
package links

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// Extract does HTTP-request GET for given URL.
// It does syntax analysis HTML and returns links from HTML-document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("retrieving %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("analysis %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // Ignoring wrong URL
				}
				links = append(links, link.String())
			}
		}
	}
	ForEachNode(doc, visitNode, nil)
	return links, nil
}

// ForEachNode calls functions pre(x) and post(x) for each node x
// in tree with root n. Both functions are not mandatory.
// pre is called before visiting child nodes, and post - after.
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// Visit adds all links, that were found in Node "n",
// to "links" slice and returns result
func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

// VisitAddTask need to be done, isn't working
func VisitAddTask(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if c := n.FirstChild; c != nil {
		links = VisitAddTask(VisitAddTask(links, c), c.NextSibling)
	}
	return links
}

// FindLinks does HTTP-request GET for given link.
// Does syntax analysis of answer as HTML-document.
// And extracts and returns links.
func FindLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("retrieving %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("analisys %s as HTML: %v", url, err)
	}
	return Visit(nil, doc), nil
}

// FindLinksLog provides logging utilities.
func FindLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return FindLinks(url)
}

/*
// CountWordsAndImages does HTTP-request GET HTML-document
// url and returns amount of words and images in it.
// P.S.: also representation of returing empty result in GoLang
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	//words, images = countWordsAndImages(doc)
	return
}


func countWordsAndImages(n *html.Node) (words, images int) {

		Write code here

}
*/

// BreadthFirst calls f for each element in worklist.
// All elements, which f returns, are added too worklist.
// f is called for each element not more then 1 time.
func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
