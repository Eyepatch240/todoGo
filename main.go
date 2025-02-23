package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"todoGo/database"
	"todoGo/handlers"
	"todoGo/models"
)

func main() {
	r := gin.Default()
	database.DB = database.InitDB()

	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.TodoList{},
		&models.Entry{},
		&models.RefreshToken{},
	); err != nil {
		log.Fatal("Failed to automigrate the db:", err)
	}

	userOps := r.Group("/users")
	{
		userOps.GET("", handlers.GetUsers)
		userOps.POST("", handlers.Signup)
		userOps.POST("/login", handlers.Login)
		userOps.DELETE("", handlers.DeleteUser)
		userOps.PUT("", handlers.EditUser)
	}

	listOps := r.Group("/todo")
	{
		listOps.POST("/get", handlers.GetLists)
		listOps.POST("", handlers.CreateList)
		listOps.DELETE("", handlers.DeleteList)
		listOps.POST("/entry/get", handlers.GetEntries)
		listOps.POST("/entry", handlers.CreateEntry)
		listOps.DELETE("/entry", handlers.DeleteEntry)
		listOps.PUT("/entry", handlers.UpdateEntry)
	}

	r.Run()
}
