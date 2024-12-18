package routes

import (
	"devops-assignment/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.POST("/create-product", controllers.CreateProduct)
	router.GET("/get-products", controllers.GetAllProducts)
	router.GET("/get-product/:id", controllers.GetProductByID)
	router.PATCH("/update-product/:id", controllers.UpdateProduct)
	router.DELETE("/delete-product/:id", controllers.DeleteProduct)
}
