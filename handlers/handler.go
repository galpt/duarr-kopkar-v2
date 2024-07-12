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

func HandleRegistrasi(c *gin.Context) {
	var (
		acc Account
	)

	c.Bind(&acc)

	templatePrint := fmt.Sprintf("Nama: %v\nEmail: %v\nUsername: %v\nPassword: %v\nNoHandphone: %v\n", acc.Nama, acc.Email, acc.Username, acc.Password, acc.NoHandphone)
	fmt.Println(templatePrint)

	c.JSON(200, gin.H{
		"Nama":        acc.Nama,
		"Email":       acc.Email,
		"Username":    acc.Username,
		"Password":    acc.Password,
		"NoHandphone": acc.NoHandphone,
	})
}

func ApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Server time: %v", time.Now().UTC().Format(time.RFC850)),
	})
}
