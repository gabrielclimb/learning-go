package router

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.ShowStudentByID)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("students/:id", controllers.DeleteStudent)
	r.PATCH("students/:id", controllers.EditStudent)
	r.GET("/:name", controllers.Hello)
	r.Run(":5000")
}
