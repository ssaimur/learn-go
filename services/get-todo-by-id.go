package services

import (
	"context"
	"learn-go/models"
	"learn-go/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary      Get a Todo by ID
// @Description  Retrieves a todo item by its unique identifier.
// @Tags         todos
// @Param        id     path      string         true  "Todo ID"
// @Success      200    {object}  models.Todo
// @Failure      400    {string}  string         "Bad request"
// @Failure      404    {string}  string         "Not found"
// @Router       /todos/{id} [get]
func GetTodoById(c *gin.Context) {
	idParam := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
	}

	filter := bson.M{"_id": objectId}

	var todo models.Todo
	err = mongo.TodoCollection.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
