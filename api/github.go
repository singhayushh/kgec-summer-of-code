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
	pulls  []User
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
	Repository string `json:"repository,omitempty"`
}

// User struct
type User struct {
	Login   string `json:"name,omitempty"`
	HTMLURL string `json:"url,omitempty"`
	Pulls   []Pull `json:"user,omitempty"`
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
	count := 0
	for index, repo := range repos {
		index = index + 1
		issues, _, err := g.client.Issues.ListByRepo(g.ctx, "dsckgec", repo, &github.IssueListByRepoOptions{State: "open", Since: startOfTime})
		if err != nil {
			continue
		}
		for _, issue := range issues {
			if issue.PullRequestLinks != nil {
				continue
			}
			count++
			newIssue := Issue{*issue.Title, *issue.HTMLURL, *issue.Comments, repo}
			g.issues = append(g.issues, newIssue)
			if count > index*2 {
				break
			}
		}
	}
}

// FetchPullStats ...
func (g *GitHubAPI) FetchPullStats() {
	for index, repo := range repos {
		index = index + 1
		pulls, _, err := g.client.PullRequests.List(g.ctx, "dsckgec", repo, &github.PullRequestListOptions{State: "all"})
		if err != nil {
			continue
		}
		for _, pull := range pulls {
			newPull := Pull{*pull.Title, *pull.HTMLURL, repo}
			flag := 0
			for i := 0; i < len(g.pulls); i++ {
				user := &g.pulls[i]
				if user.Login == *pull.User.Login {
					flag = 1
					user.addPull(newPull)
				}
			}
			if flag == 0 {
				newUser := User{}
				newUser.Login = *pull.User.Login
				newUser.HTMLURL = *pull.User.HTMLURL
				newUser.Pulls = newUser.addPull(newPull)
				g.pulls = append(g.pulls, newUser)
			}

		}
	}
}

func (user *User) addPull(pull Pull) []Pull {
	user.Pulls = append(user.Pulls, pull)
	return user.Pulls
}
