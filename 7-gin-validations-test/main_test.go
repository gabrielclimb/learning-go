package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

func setupTestsRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func insertStudent() models.Student {
	student := models.Student{
		Name: "Testing",
		RG:   "440940187",
		CPF:  "34578962455",
	}
	database.DB.Create(&student)
	return student
}

func deleteStudent(id uint) {
	var student models.Student
	database.DB.Delete(&student, id)
}

func TestCheckStatusCodeHello(t *testing.T) {
	r := setupTestsRoutes()
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

	studentMocked := insertStudent()
	defer deleteStudent(studentMocked.ID)

	r := setupTestsRoutes()

	route := "/students"

	r.GET(route, controllers.ShowAllStudents)
	req := httptest.NewRequest(http.MethodGet, route, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindStudentByID(t *testing.T) {
	var student models.Student

	database.ConnectToDatabase()
	studentMocked := insertStudent()
	defer deleteStudent(studentMocked.ID)

	r := setupTestsRoutes()
	r.GET("/students/:id", controllers.FindStudentByID)

	route := fmt.Sprintf("/students/%s", strconv.Itoa(int(studentMocked.ID)))
	req := httptest.NewRequest(http.MethodGet, route, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	json.Unmarshal(res.Body.Bytes(), &student)

	assert.Equal(t, studentMocked.Name, student.Name, "names should be the same")

}

func TestDeleteStudent(t *testing.T) {
	r := setupTestsRoutes()
	database.ConnectToDatabase()
	studentMocked := insertStudent()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	route := fmt.Sprintf("/students/%s", strconv.Itoa(int(studentMocked.ID)))

	req := httptest.NewRequest(http.MethodDelete, route, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}
