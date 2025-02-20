package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoGo/database"
	"todoGo/models"
)

func CreateUser(c *gin.Context) {
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem saving the user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully"})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "there was a problem with the Database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": users})
}

func EditUser(c *gin.Context) {
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, req.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Email = req.Email
	user.Password = req.Password

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem updating the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	if err := database.DB.Delete(&req).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "there was a problem deleting the User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
