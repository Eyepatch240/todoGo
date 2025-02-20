package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoGo/database"
	"todoGo/models"
)

func CreateUser(c *gin.Context) {
	db := database.InitDB()
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	if err := db.Save(&req).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem saving the user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully"})
}

func GetUsers(c *gin.Context) {
	db := database.InitDB()
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "there was a problem with the Database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": users})
}

func EditUser(c *gin.Context) {
	db := database.InitDB()
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	var user models.User
	if err := db.First(&user, req.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Email = req.Email
	user.Password = req.Password

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem updating the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	db := database.InitDB()
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem processing your request"})
		return
	}

	if err := db.Delete(&req).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": "there was a problem deleting the User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
