package routes

import (
	"danieljonguitud.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	userEventsGroup := server.Group("/")
	userEventsGroup.Use(middlewares.Authenticate)
	userEventsGroup.POST("/events", createEvent)
	userEventsGroup.PUT("/events/:id", updateEvent)
	userEventsGroup.DELETE("/events/:id", deleteEvent)
	userEventsGroup.POST("/events/:id/register", registerForEvent)
	userEventsGroup.DELETE("/events/:id/register", cancelRegistration)

	// Users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
