package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// usersGroup = router.Group("users")

	message := hola("jabez")
	router.GET("/ping", func(con *gin.Context) {
		con.JSON(200, gin.H{
			"message": message,
		})
	})
	router.Run()
}
