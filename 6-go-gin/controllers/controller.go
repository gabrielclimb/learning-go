package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":      "1",
		"student": "Gabriel",
	})
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API say": fmt.Sprintf("Hello %s", name),
	})
}
