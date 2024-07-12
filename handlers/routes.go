package handlers

func (server *Server) home() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) login() {
	server.Router.GET("/login", RenderLoginPage)
	server.Router.POST("/login", RenderLoginPage) // Add POST method for login
}

func (server *Server) register() {
	server.Router.GET("/register", RenderRegisterPage)
	server.Router.POST("/register", RenderRegisterPage) // Add POST method for register
}

func (server *Server) apiStatus() {
	server.Router.GET("/api", ApiStatus)
}
