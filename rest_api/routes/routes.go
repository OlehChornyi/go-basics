package routes

import (
	"example.com/rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Register handlers for incoming http requests
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", singup)
	server.POST("/login", login)
}
