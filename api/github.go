package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var repos [13]string = [13]string{"space-missions", "computer-vision-into-reality", "multipurpose-chatbot", "heart-saver", "hello-ml", "passman", "androlearn", "resumie", "cleanurge-mcu", "clearnurge-backend", "libraryly", "sac-kgec-web", "dsck-website"}

var startOfTime time.Time = time.Date(2021, time.Month(4), 10, 0, 0, 0, 0, time.UTC)

// GitHubAPI .. stores github api and context
type GitHubAPI struct {
	client *github.Client
	ctx    context.Context
	issues []Issue
	pulls  []Pull
}

// Issue struct
type Issue struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// Pull struct
type Pull struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
	User  User   `json:"user,omitempty"`
}

// User struct
type User struct {
	Login   string `json:"name,omitempty"`
	HTMLURL string `json:"url,omitempty"`
}

// Init ...
func (g *GitHubAPI) Init() {
	g.ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(g.ctx, ts)
	g.client = github.NewClient(tc)
	g.FetchPullStats()
	fmt.Println(g.pulls)
}

// FetchIssueStats ...
func (g *GitHubAPI) FetchIssueStats() {
	for _, repo := range repos {
		issues, _, err := g.client.Issues.ListByRepo(g.ctx, "dsckgec", repo, &github.IssueListByRepoOptions{State: "open", Since: startOfTime})
		if err != nil {
			return
		}
		for _, issue := range issues {
			newIssue := Issue{*issue.Title, *issue.URL}
			g.issues = append(g.issues, newIssue)
		}
	}
}

// FetchPullStats ...
func (g *GitHubAPI) FetchPullStats() {
	for _, repo := range repos {
		pulls, _, err := g.client.PullRequests.List(g.ctx, "dsckgec", repo, &github.PullRequestListOptions{State: "closed"})
		if err != nil {
			return
		}
		for _, pull := range pulls {
			newPull := Pull{*pull.Title, *pull.URL, User{*pull.User.Login, *pull.User.HTMLURL}}
			g.pulls = append(g.pulls, newPull)
		}
	}
}
