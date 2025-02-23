package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todoGo/models"
	"todoGo/service"
)

func Signup(c *gin.Context) {
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there was an error processing the json"})
		return
	}

	err := service.Signup(&req)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "")
}

func Login(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there was an error processing the json"})
		return
	}

	var tokens []string

	tokens, err := service.Login(&req)

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("accessToken", tokens[0], int((5 * time.Minute).Seconds()), "/", "localhost", false, true)

	c.SetCookie("refreshToken", tokens[1], int((15 * 24 * time.Hour).Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
