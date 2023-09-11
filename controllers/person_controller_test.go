package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/MikeMwita/person-api/database"
	"github.com/MikeMwita/person-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	// Set up a test router and database connection
	r := gin.Default()
	database.ConnectDatabase()

	// Create a test request with JSON data
	reqBody := models.Person{Name: "John Doe"}
	reqJSON, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/person", bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(recorder, req)

	// Check the response status code (200 OK for a successful request)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check the response JSON content (you can assert the response data here)
	var response models.Person
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "John Doe", response.Name)
}
