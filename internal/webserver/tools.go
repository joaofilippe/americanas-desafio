package webserver

import "github.com/gin-gonic/gin"

// GetTools is a handler to provide some tools to api
func (w *WebApp) GetTools(routes *gin.RouterGroup) {
	routes.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}