package routes

import (
	"github.com/MikeMwita/person-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//r.POST("/api", controllers.CreatePerson)
	r.POST("/api/:name", controllers.CreatePerson)
	//r.GET("/api/:name", controllers.GetPersonByName)
	r.GET("/api/:user_id", controllers.GetPersonByID)
	r.PUT("/api/:user_id", controllers.UpdatePerson)
	r.DELETE("/api/:user_id", controllers.DeletePerson)

	return r
}
