package controllers

import (
	"github.com/MikeMwita/person-api/database"
	"github.com/MikeMwita/person-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.FirstOrCreate(&person)
	c.JSON(http.StatusOK, person)
}

func GetPersonByName(c *gin.Context) {
	name := c.Param("name")
	var people []models.Person

	// Check if name is a string
	if reflect.TypeOf(name).Kind() != reflect.String {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name must be a string"})
		return
	}

	database.DB.Find(&people, "name = ?", name)
	c.JSON(http.StatusOK, people)
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Param("id")

	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}

	database.DB.Updates(&person)
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Person{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Person deleted successfully"})
}
