// backend/main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"yona-backend/controllers"
	"yona-backend/models"
	"yona-backend/repositories"
	"yona-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=yona port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Clean database
	db.Exec("DROP TABLE IF EXISTS users")
	db.AutoMigrate(&models.User{})

	return db
}

func setupRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	db := setupTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	return router, db
}

func createTestUser(t *testing.T, db *gorm.DB) models.User {
	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	result := db.Create(&user)
	assert.NoError(t, result.Error)
	return user
}

func TestGetAllUsers(t *testing.T) {
	router, db := setupRouter(t)
	user := createTestUser(t, db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, user.ID, users[0].ID)
}

func TestCreateUser(t *testing.T) {
	router, _ := setupRouter(t)

	newUser := models.User{
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}
	jsonUser, _ := json.Marshal(newUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.NotZero(t, createdUser.ID)
}

func TestCreateUserDuplicate(t *testing.T) {
	router, db := setupRouter(t)
	existingUser := createTestUser(t, db)

	jsonUser, _ := json.Marshal(existingUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUserByID(t *testing.T) {
	router, db := setupRouter(t)
	user := createTestUser(t, db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &fetchedUser)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
}

func TestGetUserByIDNotFound(t *testing.T) {
	router, _ := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateUser(t *testing.T) {
	router, db := setupRouter(t)
	user := createTestUser(t, db)

	updatedUser := models.User{
		Name:  "Updated Name",
		Email: "updated@example.com",
	}
	jsonUser, _ := json.Marshal(updatedUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resultUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &resultUser)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Name, resultUser.Name)
	assert.Equal(t, updatedUser.Email, resultUser.Email)
}

func TestDeleteUser(t *testing.T) {
	router, db := setupRouter(t)
	user := createTestUser(t, db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify user is deleted
	var count int64
	db.Model(&models.User{}).Where("id = ?", user.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
