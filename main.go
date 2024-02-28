package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Task struct represents the task model
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./task.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		due_date TEXT,
		status TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertSQL := "INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(insertSQL, task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	id, _ := result.LastInsertId()
	task.ID = int(id)

	c.JSON(http.StatusCreated, task)
}

// GetTask retrieves a task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")

	var task Task
	query := "SELECT id, title, description, due_date, status FROM tasks WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask updates a task by ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateSQL := "UPDATE tasks SET title=?, description=?, due_date=?, status=? WHERE id=?"
	_, err := db.Exec(updateSQL, task.Title, task.Description, task.DueDate, task.Status, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	deleteSQL := "DELETE FROM tasks WHERE id=?"
	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// ListTasks retrieves all tasks
func ListTasks(c *gin.Context) {
	var tasks []Task

	query := "SELECT id, title, description, due_date, status FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan tasks"})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func main() {
	r := gin.Default()

	r.POST("/tasks", CreateTask)
	r.GET("/tasks/:id", GetTask)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)
	r.GET("/tasks", ListTasks)

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
