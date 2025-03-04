package Handlers

import (
	"bytes"
	"encoding/json"
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

func TestParseCode_ValidRequest(t *testing.T) {
	router := setupRouter()

	requestBody, _ := json.Marshal(CodeRequest{Code: "mkdir"})

	req, _ := http.NewRequest("POST", "/parse", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Testing response status
	assert.Equal(t, http.StatusOK, w.Code, "Response code should be 200")
	// parse response body
	var jsonResponse JSONResponse
	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err)

	// validate response structure
	assert.NotEmpty(t, jsonResponse.Message)
	assert.NotNil(t, jsonResponse.Error)
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
