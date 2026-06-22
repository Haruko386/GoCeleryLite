package router

import (
	"GoCeleryLite/internal/server"
	"github.com/gin-gonic/gin"
)

type Router struct {
	httpHandler server.Server
}

func (r *Router) Setup(engine *gin.Engine) {
	engine.Use(func(c *gin.Context) {
		c.Header("X-API-Source", "go")
		c.Next()
	})

	v1 := engine.Group("/api/v1")
	{
		v1.POST("/tasks")       // Submit task
		v1.GET("/tasks")        // Get tasks
		v1.GET("/tasks/:id")    // Get task detail by ID
		v1.DELETE("/tasks/:id") // Cancel task by id

		v1.GET("/stats")  // Get queue stats
		v1.GET("/health") // Check queue health
	}
}
