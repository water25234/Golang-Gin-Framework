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
	apiv1user "github.com/water25234/Golang-Gin-Framework/api/v1/user"
)

type User struct {
	UserId string `json:"userId"`
}

func TestIUserGetRouter(t *testing.T) {
	user := User{
		UserId: "123",
	}

	router := gin.Default()
	router.GET("/api/v1/user/:uid", apiv1user.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "api/v1/user/"+user.UserId, nil)

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
