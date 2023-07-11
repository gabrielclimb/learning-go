package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

func SetupTestsRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckStatusCodeHello(t *testing.T) {
	r := SetupTestsRoutes()
	r.GET("/:name", controllers.Hello)
	req := httptest.NewRequest(http.MethodGet, "/gabriel", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "should be 200=200")

	mockAnswer := `{"API says:":"Hello gabriel, What's up?"}`
	bodyAns, err := io.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, mockAnswer, string(bodyAns))

}

func TestShowAllStudentsRoute(t *testing.T) {
	database.ConnectToDatabase()

	studentID := insertStudent()
	defer deleteStudent(studentID)

	r := SetupTestsRoutes()

	route := "/students"

	r.GET(route, controllers.ShowAllStudents)
	req := httptest.NewRequest(http.MethodGet, route, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func insertStudent() uint {
	student := models.Student{
		Name: "Testing",
		RG:   "440940187",
		CPF:  "34578962455",
	}
	database.DB.Create(&student)
	return student.ID
}

func deleteStudent(id uint) {
	var student models.Student
	database.DB.Delete(&student, id)
}
