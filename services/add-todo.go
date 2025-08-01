package services

import (
	"learn-go/models"
	"learn-go/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(context *gin.Context) {
	var newTodo struct {
		Item      string             `bson:"item" json:"item"`
		Completed bool               `bson:"completed" json:"completed"`
	}

	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := mongo.TodoCollection.InsertOne(context, newTodo)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		context.IndentedJSON(http.StatusInternalServerError, "failed to cast inserted ID")
		return
	}

	response := models.Todo {
		ID: id,
		Item: newTodo.Item,
		Completed: newTodo.Completed,
	}

	context.IndentedJSON(http.StatusCreated, response)
}
