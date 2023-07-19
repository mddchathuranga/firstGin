package handlers

import (
	"net/http"
	"strconv"

	"github.com/dasun/first_api/src/dtos"
	"github.com/dasun/first_api/src/models"
	"github.com/dasun/first_api/src/services"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	getDTOs := models.MapToGetDTOs(users)
	c.IndentedJSON(http.StatusOK, getDTOs)
}

func GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := services.GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	getDTO := models.MapToGetDTO(user)
	c.IndentedJSON(http.StatusOK, getDTO)

}
func CreateUsers(c *gin.Context) {
	var createDTO dtos.CreateDTO
	if err := c.BindJSON(&createDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User not created"})
		return
	}
	newUser := models.MapToUser(createDTO)
	services.CreateUser(newUser)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user created"})

}
func DeleteById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := services.DeleteUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}
func UpdateById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var updateDTO dtos.UpdateDTO
	if err := c.BindJSON(&updateDTO); err != nil {
		return
	}

	updateUser := models.MapToUserFromUpdateDTO(updateDTO)
	err := services.UpdateUserByID(id, updateUser)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user updated"})
}
