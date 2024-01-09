package webserver

import (

	"github.com/gin-gonic/gin"
	"github.com/joaofilippe/americanas-desafio/internal/interfaces"
)

// WebApp is a simple struct
type WebApp struct {
	Application interfaces.IListNodeService
}

// Server is a simple handler
func (w *WebApp) Server() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	listGroup := apiV1.Group("/list")
	{
		listGroup.GET("/", w.HelloWorld)
		listGroup.POST("/save_list", w.SaveLists)
		listGroup.GET("/merge/:id", w.MergeLists)
	}

	return r
}
