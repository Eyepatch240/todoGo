package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoGo/database"
	"todoGo/models"
)

func CreateList(c *gin.Context) {
	var req models.TodoList

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the request"})
		return
	}

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"message": "there was an error while operating on the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "we saved the todoList. All good."})
}

func GetLists(c *gin.Context) {

	var req struct {
		UserID int `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "There was an error processing the request"})
		return
	}

	var lists []models.TodoList

	if err := database.DB.Where("user_id = ?", req.UserID).Find(&lists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "there was an error fetching the Database man."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lists": lists})
}

func DeleteList(c *gin.Context) {
	var req models.TodoList

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the request dude"})
		return
	}

	if err := database.DB.Delete(&req).Error; err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"message": "there was an error while operating on the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "we deleted the list man, all good"})
}
