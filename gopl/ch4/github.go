package github

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchReault struct {
	TotalCount int `json: "total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json: "html_url"`
	Title string
	State string
	User *User
	CreatedAt time.time `json: "created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json: "html_url"`
}

func SearchIssue(terms []string) (*IssueSearchReault, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL+"?q="+q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.StatusCode)
	}

	var result IssueSearchReault
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}