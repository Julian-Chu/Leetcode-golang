package main

import (
	"fmt"
)

func main() {

	issue := Issue{1, "test"}
	issues := make([]*Issue, 0)
	issues = append(issues, &issue)

	var items []Issue
	items = append(items, issue)
	fmt.Println(issue)
	fmt.Println(issues)
	fmt.Println(items)

	issueSearchResult := IssueSearchResult{1, nil, nil}
	issueSearchResult.ItemPointers = issues

	fmt.Println(cap(issues))

}

type IssueSearchResult struct {
	TotalCount   int `json:"total_count"`
	ItemPointers []*Issue
	Items        []Issue
}

type Issue struct {
	Number int
	Title  string
}
