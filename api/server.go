package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server struct to store the gin Engine
type Server struct {
	Router *gin.Engine
}

// Run the server on dedicated port
func (server *Server) Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	server.initRoutes()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("views/*.html")
	server.Router.Static("/css", "views/css")
	server.Router.Static("/img", "views/img")
	server.Router.Static("/js", "views/js")

	fmt.Printf(`KSOC server running on http://localhost:%s!`, port)
	fmt.Println()
	server.Router.Run(":" + port)
}
