package main

import (
	"Book_exercises/chapter5/links"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		title(url)
	}
}

// Instead of controlling everywhere resp.Body.Close
// we just type it with "defer" prefix and guarantee the release of resources
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Check if header "Content-Type" is HTML or not.
	// (for example, "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		//resp.Body.Close()
		return fmt.Errorf("%s has type of %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	//resp.Body.Close()
	if err != nil {
		return fmt.Errorf("analysis %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	links.ForEachNode(doc, visitNode, nil)
	return nil
}
