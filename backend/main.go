// backend/main.go
package main

import (
	"yona-backend/controllers"
	"yona-backend/repositories"
	"yona-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	router := gin.Default()

	// User routes
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Start server
	router.Run(":8080")
}
