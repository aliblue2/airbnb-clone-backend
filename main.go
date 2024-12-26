package main

import (
	"airbnb.com/airbnb/db"
	"airbnb.com/airbnb/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RouterHandler(server)

	server.Run(":8080")

}
