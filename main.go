package main

import (
	"github.com/juliasilvamoura/gin-api-rest/database"
	"github.com/juliasilvamoura/gin-api-rest/routes"
)

func main() {
	database.ConectDatabase()
	routes.HandleRequests()
}
