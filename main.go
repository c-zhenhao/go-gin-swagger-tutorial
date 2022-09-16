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
	router.GET("/todo/:id", getTodoById)
	router.POST("/todo", createTodo)
	router.DELETE("/todo/:id", deleteTodo)
	router.PATCH("/todo/:id", updateTodo)

	// run gin server
	router.Run("localhost:8000")
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

func getTodoById(c *gin.Context) {
	// capture params from the request
	id := c.Param("id")

	// loop through todolist and return item with the ID
	for _, todo := range todoList {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	// return error if not found
	c.JSON(http.StatusNotFound, Message{Message: "Todo not found"})
}

func createTodo(c *gin.Context) {
	// capture request body
	var newTodo Todo

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&newTodo); err != nil {
		r := Message{"an error has occurred while creating todo"}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	// add the new todo to todoList
	todoList = append(todoList, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	// capture params from request
	id := c.Param("id")

	// loop through todoList and delete item with the ID
	for index, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:index], todoList[index+1:]...)
			r := Message{"Todo deleted successfully"}
			c.JSON(http.StatusOK, r)
			return
		}
	}

	// return error message if todo not found
	r := Message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}

func updateTodo(c *gin.Context) {
	// capture params from request
	id := c.Param("id")

	// loop through todoList and update item with the ID
	for index, todo := range todoList {
		if todo.ID == id {
			// capture request body
			var updatedTodo Todo

			// bind the received JSON data to updatedTodo
			if err := c.BindJSON(&updatedTodo); err != nil {
				r := Message{"an error has occurred while updating todo"}
				c.JSON(http.StatusBadRequest, r)
				return
			}

			// update the todo
			todoList[index] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	// return error message if todo not found
	r := Message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}
