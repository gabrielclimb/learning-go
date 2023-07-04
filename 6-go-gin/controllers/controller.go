package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":      "1",
		"student": "Gabriel",
	})
}
