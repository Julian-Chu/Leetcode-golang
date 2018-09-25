package main

func main() {

	issue := Issue{1, "test"}
	issues := make([]Issue, 2)
	issueSearchResult := IssueSearchResult{1, new}
	issueSearchResult.Items = &issues

}

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number int
	Title  string
}
