package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Ini untuk handle pathname (istilah lainnya "routing")
func RenderHome(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func RenderLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func RenderRegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func ApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server time: %v", time.Now().UTC().Format(time.RFC850)),
	})
}
