package handlers

func (server *Server) home() {
	// ini untuk munculkan halaman homepage
	server.Router.GET("/", RenderHome)
}

func (server *Server) login() {
	// ini untuk munculkan halaman login
	server.Router.GET("/login", RenderLoginPage)
<<<<<<< Updated upstream
	server.Router.POST("/login", RenderLoginPage) // Add POST method for login
=======

	// ini untuk handle login ketika submit form dengan method POST
	server.Router.POST("/login", RenderLoginPage)
>>>>>>> Stashed changes
}

func (server *Server) register() {
	// ini untuk munculkan halaman registrasi
	server.Router.GET("/register", RenderRegisterPage)
<<<<<<< Updated upstream
	server.Router.POST("/register", RenderRegisterPage) // Add POST method for register
=======

	// ini untuk handle registrasi ketika submit form dengan method POST
	server.Router.POST("/register", HandleRegistrasi)
>>>>>>> Stashed changes
}

func (server *Server) apiStatus() {
	// ini untuk munculkan status API
	server.Router.GET("/api", ApiStatus)
}
