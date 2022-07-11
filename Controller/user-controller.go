package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/Models"
)
//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User

	if err := c.ShouldBindJSON(&user) ; err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		err := Models.CreateUser(&user)
		if err != nil {
			//fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}

}
//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Message":  user.Name + " is deleted"})
		Models.DeleteUser(&user, id)
	}
}

func PatchUser(c *gin.Context){
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)

	if err != nil {
		c.JSON(http.StatusNotFound, user)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}