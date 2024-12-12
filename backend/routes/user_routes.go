package routes

import (
	"yona-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
