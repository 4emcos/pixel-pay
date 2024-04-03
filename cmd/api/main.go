package main

import (
	"github.com/gin-gonic/gin"
	"pixel-pay/cmd/api/routes"
	"pixel-pay/database"
)

func main() {

	db := database.Connect()

	router := routes.NewRouter(gin.Default(), db)
	router.SetupRoutes()

	router.Run(":8080")
}
