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
