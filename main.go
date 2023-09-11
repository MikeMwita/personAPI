package main

import (
	"github.com/MikeMwita/person-api/database"
	"github.com/MikeMwita/person-api/routes"
)

func main() {
	database.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8000")
}
