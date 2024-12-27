package main

import (
	"devops-assignment/config"
	"devops-assignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.UserRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":3005")
}
