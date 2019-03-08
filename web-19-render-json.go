package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Desc string `json:"desc"`
	Done bool   `json:"done"`
}

func main() {
	r := gin.Default()

	r.GET("/tasks/", func(c *gin.Context) {
		tasks := []Task{
			{Desc: "Learn Go"},
			{Desc: "Learn Gin"},
		}
		c.JSON(http.StatusOK, tasks)
	})
	r.Run(":8000")
}
