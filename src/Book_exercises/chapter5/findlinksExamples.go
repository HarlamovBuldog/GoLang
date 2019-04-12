package main

import (
	"Book_exercises/chapter5/links"
	"fmt"
	"log"
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
		for _, link := range links.Visit(nil, doc) {
			fmt.Println(link)
		}
	*/
	for _, link2 := range links.VisitAddTask(nil, doc) {
		fmt.Println(link2)
	}
	for _, url := range os.Args[1:] {
		links, err := links.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

	links.BreadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
