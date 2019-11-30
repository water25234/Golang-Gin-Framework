package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"../router"
	"github.com/stretchr/testify/assert"
)

type User struct {
	UserId string `json:"userId"`
}

func TestIUserGetRouter(t *testing.T) {
	user := User{
		UserId: "123",
	}
	expectedBody, _ := json.Marshal(user)
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+user.UserId, nil)

	router.ServeHTTP(w, req)

	input := string(w.Body.String())

	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedBody), input)
}
