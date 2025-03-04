package Handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/parse", ParseCodeHandler)
	return r
}

func TestParseCode_ValirRequest(t *testing.T) {
	router := setupRouter()

	jsonBody := `{"code": "mkdir ola"}`

	req, _ := http.NewRequest("POST", "/parse", bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expectedResponde := `{"message": "Code received", "code": "mkdir ola"}`

	assert.Equal(t, http.StatusOK, w.Code, "Response code should be 200")
	assert.JSONEq(t, expectedResponde, w.Body.String(), "Response body should be equal")
}

func TestParseCode_InvalidRequest(t *testing.T) {
	router := setupRouter()

	invalidJsonBody := `{"string": "mkdir ola"}`

	req, _ := http.NewRequest("POST", "/parse", bytes.NewBuffer([]byte(invalidJsonBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Response code should be 400")
	assert.Contains(t, w.Body.String(), "Invalid JSON format", "Response body should contain error message")
}
