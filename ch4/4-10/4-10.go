package main

import (
	"fmt"
	"gopl-practise/ch4/4-10/github"
	"log"
	"os"
	"time"
)

const (
	lessOneMonth string = "less than a month"
	lessOneYear  string = "less than a year"
	moreOneYear  string = "more than a year"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	issueCount := make(map[string][]github.Issue, 3)

	for _, item := range result.Items {
		item := *item
		Y, M, _ := item.CreatedAt.Date()
		curY, curM, _ := time.Now().Date()
		switch {
		case curY-Y > 1:
			issueCount[moreOneYear] = append(issueCount[moreOneYear], item)
		case curM-M > time.Month(1):
			issueCount[lessOneYear] = append(issueCount[lessOneYear], item)
		case curM-M <= time.Month(1):
			issueCount[lessOneMonth] = append(issueCount[lessOneMonth], item)
		}
	}
	var total int
	for class, issues := range issueCount {
		fmt.Printf("class: %s, issues: %d\n", class, len(issues))
		total += len(issues)
	}
	fmt.Printf("Total: %d", total)
}
