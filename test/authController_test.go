package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	apiv1auth "github.com/water25234/Golang-Gin-Framework/api/v1/auth"
)

type Auth struct {
	ThrottleCount int
	UserID        string `json:"userId"`
}

func TestGetAuthRouter(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/auth", apiv1auth.GetAuth)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "api/v1/auth", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, Justin Home!", w.Body.String())
}

func TestAuthDeleteAuthRouter(t *testing.T) {
	id := "123"

	router := gin.Default()
	router.DELETE("/api/v1/auth/:id", apiv1auth.DeleteAuth)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "api/v1/auth/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World DELETE Justin "+id, w.Body.String())
}

func TestAuthPostAuthRouter(t *testing.T) {
	auth := Auth{
		ThrottleCount: 1,
		UserID:        "123",
	}

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, router := gin.CreateTestContext(resp)

	router.Use(func(c *gin.Context) {
		c.Set("ThrottleCount", 1)
	})

	router.POST("/api/v1/auth/:uid", apiv1auth.PostAuth)

	w := httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodPost, "api/v1/auth/"+auth.UserID, nil)

	router.ServeHTTP(w, c.Request)

	input := string(w.Body.String())

	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, GetSuccessResponse1(auth), input)
}

func GetSuccessResponse1(auth Auth) string {

	type Response struct {
		Data     Auth              `json:"data"`
		Metadata map[string]string `json:"metadata"`
	}

	rsponse := Response{
		Metadata: map[string]string{
			"status": "0000",
			"desc":   "success",
		},
		Data: auth,
	}
	b, err := json.Marshal(rsponse)

	if err != nil {
		fmt.Println("error:", err)
	}

	return string(b)
}
