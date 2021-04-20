package api

import (
	"github.com/gin-gonic/gin"
)

// RenderHome ...
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "KGEC Summer of Code 2021",
		"isHome": true,
	})
}

//dashboard
func RenderDash(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{
		"title":  "Dashboard",
		"isHome": false,
	})
}

//projects
func RenderProj(c *gin.Context) {
	c.HTML(200, "projects.html", gin.H{
		"title":  "Projects",
		"isHome": false,
	})
}

//about
func RenderAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title":  "About",
		"isHome": false,
	})
}
