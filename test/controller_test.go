package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../router"
	"github.com/stretchr/testify/assert"
)

func TestGetRouter(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, Justin Home!", w.Body.String())
}

func TestIHelloDeleteRouter(t *testing.T) {
	id := "123"
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/hello/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World DELETE Justin "+id, w.Body.String())
}
