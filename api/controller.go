package api

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

// RenderHome ...
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "KGEC Summer of Code 2021",
		"isHome": true,
	})
}

// RenderGit ...
func RenderGit(c *gin.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Issues.ListByRepo(ctx, "dsckgec", "community-discord-bot", nil)
	if err == nil {
		c.HTML(200, "dashboard.html", gin.H{
			"title": "KGEC Summer of Code 2021",
			"json":  repos,
		})
	} else {
		c.HTML(200, "dashboard.html", gin.H{
			"title": "KGEC Summer of Code 2021",
			"json":  err,
		})
	}
}
