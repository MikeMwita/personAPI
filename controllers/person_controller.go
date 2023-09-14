package controllers

import (
	"github.com/MikeMwita/person-api/database"
	"github.com/MikeMwita/person-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePerson(c *gin.Context) {
	name := c.Param("name")

	person := models.Person{Name: name}

	if err := database.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	// Respond with the created person
	c.JSON(http.StatusOK, person)
}

func GetPersonByID(c *gin.Context) {
	id := c.Param("id")

	// Query the database to find the person by ID
	var person models.Person
	if err := database.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

//func UpdatePerson(c *gin.Context) {
//	var person models.Person
//	id := c.Param("id")
//
//	if err := database.DB.First(&person, id).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
//		return
//	}
//
//	database.DB.Updates(&person)
//	c.JSON(http.StatusOK, person)
//}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	userID := c.Param("user_id")

	if err := database.DB.First(&person, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the person in the database
	database.DB.Save(&person)

	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	userID := c.Param("user_id")

	// Delete the person with the provided 'user_id' from the database
	if err := database.DB.Where("id = ?", userID).Delete(&models.Person{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
