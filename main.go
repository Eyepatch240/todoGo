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
	db := database.InitDB()

	if err := db.AutoMigrate(
		&models.User{},
		&models.TodoList{},
		&models.Entry{},
	); err != nil {
		log.Fatal("Failed to automigrate the db:", err)
	}

	userOps := r.Group("/users")
	{
		userOps.GET("", handlers.GetUsers)
		userOps.POST("", handlers.CreateUser)
		userOps.DELETE("", handlers.DeleteUser)
		userOps.PUT("", handlers.EditUser)
	}
	r.Run()
}
