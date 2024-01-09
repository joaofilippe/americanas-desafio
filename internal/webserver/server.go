package webserver

import (

	"github.com/gin-gonic/gin"
	"github.com/joaofilippe/americanas-desafio/internal/interfaces"
)

// WebApp is a simple struct
type WebApp struct {
	Application interfaces.IListNodeService
}

// Server runs the webserver
func (w *WebApp) Server() *gin.Engine {
	r := gin.Default()
	
	apiV1 := r.Group("/api/v1")
	
	w.GetTools(apiV1)
	w.AddListGroupRoutes(apiV1, "/list")

	return r
}
