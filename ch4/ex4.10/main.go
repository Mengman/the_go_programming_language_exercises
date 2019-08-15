package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	IssuesURL    = "https://api.github.com/search/issues"
	MonthSeconds = 30 * 24 * 3600
	YearSeconds  = 12 * MonthSeconds
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	var issueMonth, issueYear, issueMoreThanYear []*Issue
	now := time.Now()
	for _, item := range result.Items {
		diff := int64(now.Sub(item.CreatedAt).Seconds())

		if diff < MonthSeconds {
			issueMonth = append(issueMonth, item)
		} else if diff < YearSeconds {
			issueYear = append(issueYear, item)
		} else {
			issueMoreThanYear = append(issueMoreThanYear, item)
		}
	}

	fmt.Printf("\nissue less than a month\n")
	for _, item := range issueMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nissue less than a year\n")
	for _, item := range issueYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nissue more than a year\n")
	for _, item := range issueMoreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}
