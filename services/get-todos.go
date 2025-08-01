package services

import (
	"learn-go/models"
	"learn-go/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func GetTodos(context *gin.Context) {
	cursor, err := mongo.TodoCollection.Find(context, bson.M{})
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context)

	var todos []models.Todo
	if err := cursor.All(context, &todos); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, todos)
}
