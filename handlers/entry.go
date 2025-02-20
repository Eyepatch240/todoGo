package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoGo/database"
	"todoGo/models"
)

func CreateEntry(c *gin.Context) {
	var req models.Entry

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the request man"})
		return
	}

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save the entry"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "all good man, we saved the entry."})
}

func GetEntries(c *gin.Context) {
	var req struct {
		ListID int `json:"todo_list_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "There was an error processing the request"})
		return
	}

	var entries []models.Entry

	if err := database.DB.Where("todo_list_id = ?", req.ListID).Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch entries for the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entries": entries})
}

func DeleteEntry(c *gin.Context) {
	var req models.Entry

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the request man"})
		return
	}

	if err := database.DB.Delete(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete the entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "the entry was deleted accordingly."})
}

func UpdateEntry(c *gin.Context) {
	var req models.Entry

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the request man"})
		return
	}

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update the entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "We managed to update the entry man, all good"})
}
