package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
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
	ans := httptest.NewRecorder()
	r.ServeHTTP(ans, req)
	assert.Equal(t, http.StatusOK, ans.Code, "should be 200=200")

	mockAnswer := `{"API says:":"Hello gabriel, What's up?"}`
	bodyAns, err := io.ReadAll(ans.Body)

	assert.NoError(t, err)
	assert.Equal(t, mockAnswer, string(bodyAns))

}
