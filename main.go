package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo represents data about a task in the todo list
type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type Message struct {
	Message string `json:"messasge"`
}

// todo slice to seed todo list data
var todoList = []Todo{
	{ID: "1", Task: "Learn Go"},
	{ID: "2", Task: "Build an API with Go"},
	{ID: "3", Task: "Document API with swagger"},
}

func main() {
	// configure the gin server with default
	router := gin.Default()

	// regiseter getAllTodos handler to the Gin router
	router.GET("/todo", getAllTodos)

	// run gin server
	router.Run("localhost:8000")
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}
