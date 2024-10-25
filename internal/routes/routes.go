package routes

import (
	"go_crud/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, personHandler *handler.PersonHandler) {
	r.POST("/person", personHandler.CreatePerson)
	r.GET("/person", personHandler.GetAllPersons)
	r.GET("/person/:id", personHandler.GetPersonByID)
	r.PUT("/person/:id", personHandler.UpdatePerson)
	r.DELETE("/person/:id", personHandler.DeletePerson)

	// Requests to non-existing endpoints
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Resource not found"})
	})
}
