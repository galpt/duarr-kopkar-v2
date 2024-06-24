package handlers

func (server *Server) home() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) apiStatus() {
	server.Router.GET("/api", ApiStatus)
}
