package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Desc string `json:"desc"`
	Done bool   `json:"done"`
}

func main() {
	r := gin.Default()

	r.POST("/tasks/", func(c *gin.Context) {
		var t Task
		if err := c.Bind(&t); err != nil {
			return
		}
		log.Printf("%+v\n", t)
	})
	r.Run(":8000")
}
