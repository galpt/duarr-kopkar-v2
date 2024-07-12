package handlers

func (server *Server) home() {
	// ini untuk munculkan halaman homepage
	server.Router.GET("/", RenderHome)
}

func (server *Server) login() {
	// ini untuk munculkan halaman login
	server.Router.GET("/login", RenderLoginPage)

	// ini untuk handle login ketika submit form dengan method POST
	server.Router.POST("/login", RenderLoginPage)
}

func (server *Server) register() {
	// ini untuk munculkan halaman registrasi
	server.Router.GET("/register", RenderRegisterPage)

	// ini untuk handle registrasi ketika submit form dengan method POST
	server.Router.POST("/register", HandleRegistrasi)
}

func (server *Server) apiStatus() {
	// ini untuk munculkan status API
	server.Router.GET("/api", ApiStatus)
}
