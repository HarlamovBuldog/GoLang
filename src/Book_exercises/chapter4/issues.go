// Prints table - search result in GitHub
// Task 4.10 realization from book
package main

import (
	"Book_exercises/chapter4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	categories := make(map[string][]*github.Issue)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		daysPassedSinceNow := time.Since(item.CreatedAt).Hours() / 24
		if daysPassedSinceNow < 31 {
			categories["lessThenMonthAgo"] = append(categories["lessThenMonthAgo"], item)
		} else if daysPassedSinceNow < 365 {
			categories["lessThenYearAgo"] = append(categories["lessThenYearAgo"], item)
		} else {
			categories["moreThenYearAgo"] = append(categories["moreThenYearAgo"], item)
		}
	}
	fmt.Printf("%d topics:\n", result.TotalCount)
	for category, itemArray := range categories {
		fmt.Printf("%d posted %s:\n", len(itemArray), category)
		fmt.Println("--------------------------------------------")
		for _, item := range itemArray {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
		fmt.Println("--------------------------------------------")
	}
}
