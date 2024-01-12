package main

import (
	"repositoryPattern/config"
	"repositoryPattern/modules/auth"
	"repositoryPattern/modules/test"

	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("api/v1")
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))

	test.NewHandler(v1, test.TestRegistry(db))

	gin.SetMode(gin.ReleaseMode)

	port := "99" // Port default jika tidak ada ASPNETCORE_PORT

	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	router.Run(":" + port)
}
