package handlers

func (server *Server) home() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) login() {
	server.Router.GET("/login", RenderLoginPage)
}

func (server *Server) apiStatus() {
	server.Router.GET("/api", ApiStatus)
}
