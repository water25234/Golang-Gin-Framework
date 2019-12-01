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

type User struct {
	UserId string `json:"userId"`
}

func TestIUserGetRouter(t *testing.T) {
	user := User{
		UserId: "123",
	}
	//expectedBody, _ := json.Marshal(user)
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+user.UserId, nil)

	router.ServeHTTP(w, req)

	input := string(w.Body.String())

	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, GetSuccessResponse(user), input)
}

func GetSuccessResponse(data User) string {

	type Response struct {
		Data     User              `json:"data"`
		Metadata map[string]string `json:"metadata"`
	}

	rsponse := Response{
		Metadata: map[string]string{
			"status": "0000",
			"desc":   "success",
		},
		Data: data,
	}
	b, err := json.Marshal(rsponse)

	if err != nil {
		fmt.Println("error:", err)
	}

	return string(b)
}
