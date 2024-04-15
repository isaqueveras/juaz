package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/v1/user/:id", getUser)
	router.GET("/v1/biometry/search", search_biometry)

	router.Run(":8181")
}
