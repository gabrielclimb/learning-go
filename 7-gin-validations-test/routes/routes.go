package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:nome", controllers.Hello)
	r.POST("/students", controllers.AddNewStudent)
	r.GET("/students/:id", controllers.FindStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run(":5000")
}
