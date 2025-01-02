package main

import (
	"crud_project/db"
	"crud_project/routes"
	"crud_project/logger"
	"github.com/gin-gonic/gin"	
	
)

func main(){
	logger.Writelogs()
	logger.InfoLogger.Print("Starting application")
	db.InitDB()
	logger.InfoLogger.Print("DB Initiated")

	server := gin.Default() // returns a pointer to the gin engine

	routes.RegisterRoutes(server)
	
	server.Run(":8080") // port 
	
}
