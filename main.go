package main

import (
	"dsviewer/hola"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// usersGroup = router.Group("users")

	message := hola.Hola("jabez")
	router.GET("/ping", func(con *gin.Context) {
		con.JSON(200, gin.H{
			"message": message,
		})
	})
	router.Run()
}
