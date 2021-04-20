package api

import (
	"github.com/gin-gonic/gin"
)

// RenderHome ...
func (g *GitHubAPI) RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "KGEC Summer of Code 2021",
		"isHome": true,
	})
}

// RenderDashboard ...
func (g *GitHubAPI) RenderDashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{
		"title":  "Dashboard | KGEC Summer of Code 2021",
		"issues": g.issues,
		"pulls":  g.pulls,
	})
}

// Refresh ...
func (g *GitHubAPI) Refresh(c *gin.Context) {
	g.issues = g.issues[:0]
	g.pulls = g.pulls[:0]

	g.FetchIssueStats()
	g.FetchPullStats()
	c.JSON(200, gin.H{
		"message": "done",
	})
}

// RenderProject ...
func (g *GitHubAPI) RenderProject(c *gin.Context) {
	c.HTML(200, "projects.html", gin.H{
		"title":  "Projects",
		"isHome": false,
	})
}

// RenderAbout ...
func (g *GitHubAPI) RenderAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title":  "About",
		"isHome": false,
	})
}
