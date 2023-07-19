package main

import (
	"github.com/dasun/first_api/src/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users/getAllUsers", handlers.GetAllUsers)
	router.GET("/users/getUserById/:id", handlers.GetById)
	router.POST("/users/createUser", handlers.CreateUsers)
	router.DELETE("/users/deleteUserById/:id", handlers.DeleteById)
	router.PUT("/users/updateUserById/:id", handlers.UpdateById)
	router.Run("localhost:8080")
}
