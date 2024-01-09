package webserver

import (
	"github.com/gin-gonic/gin"
)

// AddListGroupRoutes adds a group of routes
func (w *WebApp) AddListGroupRoutes(apiGroup *gin.RouterGroup, path string) {
	listGroup := apiGroup.Group(path)
	{
		listGroup.POST("/save_list", w.SaveLists)
		listGroup.GET("/merge/:id", w.MergeLists)
	}
}
