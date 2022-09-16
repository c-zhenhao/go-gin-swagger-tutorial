package main

import (
	"net/http"

	_ "go-gin-swagger-tutorial/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	// configure the gin server with default
	router := gin.Default()

	// register getAllTodos handler to the Gin router
	router.GET("/todo", getAllTodos)
	router.GET("/todo/:id", getTodoById)
	router.POST("/todo", createTodo)
	router.PATCH("/todo/:id", updateTodo)
	router.DELETE("/todo/:id", deleteTodo)

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run gin server
	router.Run("localhost:8000")
}

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} Todo
// @Router /todo [get]
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} Message
// @Router /todo/{id} [get]
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

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body Todo true "Todo data"
// @Success 200 {object} Todo
// @Failure 400 {object} Message
// @Router /todo [post]
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

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} Message
// @Router /todo/{id} [delete]
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

// @Summary update a todo item by ID
// @ID update-todo-by-id
// @Produce json
// @Param id path string true "Todo ID"
// @Param data body Todo true "Todo data"
// @Success 200 {object} Todo
// @Failure 404 {object} Message
// @Router /todo/{id} [patch]
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
