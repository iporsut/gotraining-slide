package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/helloform/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK,
			`<form method="POST" action="/hello/">Name: <input name="name"/></form>`)
	})
	r.POST("/hello/", func(c *gin.Context) {
		name := c.PostForm("name")

		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello " + name,
		})
	})
	r.Run(":8000")
}
