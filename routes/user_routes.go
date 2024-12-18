package routes

import (
	"devops-assignment/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/create-user", controllers.CreateUser)
	router.GET("/get-users", controllers.GetAllUsers)
	router.GET("/get-user/:id", controllers.GetUserByID)
	router.PUT("/update-user/:id", controllers.UpdateUser)
	router.DELETE("/delete-user/:id", controllers.DeleteUser)
}
