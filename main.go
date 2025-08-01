package main

import (
	"learn-go/mongo"
	"learn-go/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main () {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongo.InitMongo()

	router := gin.Default()
	router.GET("/todos", services.GetTodos)
	router.GET("/todos/:id", services.GetTodoById)
	router.PATCH("/todos/:id", services.UpdateTodoById)
	router.POST("/todos", services.AddTodo)
	router.Run("localhost:3000")
}