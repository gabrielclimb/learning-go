package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, &students)
}

func ShowStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": fmt.Sprintf("student id %s not found", id),
		})
	} else {
		c.JSON(http.StatusOK, student)
	}

}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	return
	}
	
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)

}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API say": fmt.Sprintf("Hello %s", name),
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "student deleted"})

}
