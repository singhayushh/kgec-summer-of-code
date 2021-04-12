package api

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}
