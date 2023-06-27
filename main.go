package main

import (
	"repositoryPattern/config"
	"repositoryPattern/modules/auth"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("api/v1")
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))

	router.Run(":1000")
}
