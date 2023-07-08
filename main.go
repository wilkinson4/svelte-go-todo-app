package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodos)
	router.DELETE("/todos/:id", deleteTodoByID)
	router.PUT("/todos/:id", updateTodoByID)

	router.Run("localhost:8080")
}

type todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var todos = []todo{
	{ID: "1", Description: "Wash the dishes", Done: false},
	{ID: "2", Description: "Wash the car", Done: false},
	{ID: "3", Description: "Clean the room", Done: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodos(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func deleteTodoByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range todos {
		if a.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func updateTodoByID(c *gin.Context) {
	id := c.Param("id")

	var updatedTodo todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
		return
	}

	var index = slices.IndexFunc(todos, func(a todo) bool {
		println("id: " + id)
		println("a.ID " + a.ID)
		return a.ID == id
	})

	if index == -1 {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todos[index] = updatedTodo

	c.IndentedJSON(http.StatusOK, gin.H{"message": "todo updated"})
}
