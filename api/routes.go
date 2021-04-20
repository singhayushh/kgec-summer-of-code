package api

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
	server.Router.GET("/dashboard", RenderDash)
	server.Router.GET("/projects", RenderProj)
	server.Router.GET("/about", RenderAbout)
}
