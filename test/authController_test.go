package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"../router"
	"github.com/stretchr/testify/assert"
)

type Auth struct {
	UserID string `json:"userId"`
}

func TestGetAuthRouter(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/auth", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, Justin Home!", w.Body.String())
}

func TestAuthDeleteAuthRouter(t *testing.T) {
	id := "123"
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/auth/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World DELETE Justin "+id, w.Body.String())
}

func TestAuthPostAuthRouter(t *testing.T) {
	auth := Auth{
		UserID: "123",
	}
	// expectedBody, _ := json.Marshal(Auth)
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/auth/"+auth.UserID, nil)

	router.ServeHTTP(w, req)

	input := string(w.Body.String())

	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, GetSuccessResponse_1(auth), input)
}

func GetSuccessResponse_1(auth Auth) string {

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
