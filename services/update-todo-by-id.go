package services

import (
	"context"
	"learn-go/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTodoDto struct { Completed bool `json:"completed"` } 

// @Summary Update todo
// @Description Update a todo by it's id
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body UpdateTodoDto true "Is completed"
// @Param id path string true "Todo ID"
// @Success 201 {object} models.Todo
// @Failure 400 {string} string "Bad Request"
// @Router /todos/{id} [patch]
func UpdateTodoById(c *gin.Context) {
	idParam := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
	}

	

	var body struct { Completed bool `json:"completed"` }
	err = c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": body.Completed}}

	_, err = mongo.TodoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "done")
}
