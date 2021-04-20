package api

import (
	"context"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var repos [14]string = [14]string{"space-missions", "computer-vision-into-reality", "multipurpose-chatbot", "heart-saver", "hello-ml", "passman", "androlearn", "resumie", "cleanurge-mcu", "cleanurge-backend", "libraryly", "sac-kgec-web", "dsck-website", "cleanurge-app"}

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
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
	Comments   int    `json:"comments,omitempty"`
	Repository string `json:"repository,omitempty"`
}

// Pull struct
type Pull struct {
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
	User       User   `json:"user,omitempty"`
	Repository string `json:"repository,omitempty"`
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
	g.FetchIssueStats()
	g.FetchPullStats()
}

// FetchIssueStats ...
func (g *GitHubAPI) FetchIssueStats() {
	for _, repo := range repos {
		issues, _, err := g.client.Issues.ListByRepo(g.ctx, "dsckgec", repo, &github.IssueListByRepoOptions{State: "open", Since: startOfTime})
		if err != nil {
			continue
		}
		for _, issue := range issues {
			newIssue := Issue{*issue.Title, *issue.HTMLURL, *issue.Comments, repo}
			g.issues = append(g.issues, newIssue)
		}
	}
}

// FetchPullStats ...
func (g *GitHubAPI) FetchPullStats() {
	for _, repo := range repos {
		pulls, _, err := g.client.PullRequests.List(g.ctx, "dsckgec", repo, &github.PullRequestListOptions{State: "closed"})
		if err != nil {
			continue
		}
		for _, pull := range pulls {
			newPull := Pull{*pull.Title, *pull.HTMLURL, User{*pull.User.Login, *pull.User.HTMLURL}, repo}
			g.pulls = append(g.pulls, newPull)
		}
	}
}
