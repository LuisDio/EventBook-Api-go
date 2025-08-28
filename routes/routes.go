package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)       // should be protected
	server.PUT("/events/:id", updateEvent)    // should be protected
	server.DELETE("/events/:id", deleteEvent) // should be protected
	server.POST("/signup", signup)
	server.POST("/login", login)
}
