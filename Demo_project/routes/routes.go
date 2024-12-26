package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// register handlers for incoming http request.
	server.GET("/events", GetEvents) 
	server.POST("/events", CreateEvent)
	server.GET("/events/:id", GetEvent) 
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent) 
}