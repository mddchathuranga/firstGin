package main

import (
	//"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

var users = []user{
	{ID: 1, Username: "Dasun", Address: "Mirigama", Mobile: "0711234567", Age: 30, Email: "dasun.e.com"},
	{ID: 2, Username: "Amith", Address: "Colombo", Mobile: "0722222222", Age: 15, Email: "amith.e.com"},
	{ID: 3, Username: "Lasan", Address: "Gampaha", Mobile: "07188888888", Age: 25, Email: "lasan.e.com"},
}

/*func getUserById(id int) (*user, error) {
	for i, u := range users {
		if u.ID == id {

		}
	}
	return nil, errors.New("user not found")
}*/

//user handlers

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, users[i])
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

}

func createUsers(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func deleteById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "user deleted"})
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return

	}

}
func updateById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, u := range users {
		if u.ID == id {
			var updateUser user
			if err := c.BindJSON(&updateUser); err != nil {
				return

			}
			users[i].Username = updateUser.Username
			users[i].Address = updateUser.Address
			users[i].Mobile = updateUser.Mobile
			users[i].Age = updateUser.Age
			users[i].Email = updateUser.Email

			c.IndentedJSON(http.StatusOK, gin.H{"message": "user updated"})
			return

		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	return

}

func main() {

	router := gin.Default()
	router.GET("/users/getAllUsers", getAllUsers)
	router.GET("/users/getUserById/:id", getById)
	router.POST("/users/createUser", createUsers)
	router.DELETE("/users/deleteUserById/:id", deleteById)
	router.PUT("/users/updateUserById/:id", updateById)
	router.Run("localhost:8080")

}
