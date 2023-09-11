package routes

import (
	"github.com/MikeMwita/person-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api", controllers.CreatePerson)
	r.GET("/api/:name", controllers.GetPersonByName)
	r.PUT("/api/:id", controllers.UpdatePerson)
	r.DELETE("/api/:id", controllers.DeletePerson)
	return r
}
