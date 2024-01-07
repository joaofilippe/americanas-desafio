package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaofilippe/americanas-desafio/api/requests"
	"github.com/joaofilippe/americanas-desafio/interfaces"
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

// HelloWorld is a simple handler
func (w *WebApp) HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// SaveLists is a simple handler
func (w *WebApp) SaveLists(c *gin.Context) {
	l := new(requests.SaveListsRequest)

	c.Bind(l)

	id, err := w.Application.SaveListsNode(l.List1, l.List2)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "it works",
		"data":    id,
		"lists":   l,
	})
}

// MergeLists is a simple handler
func (w *WebApp) MergeLists(c *gin.Context) {
	id := c.Param("id")

	strconv.Atoi(id)

	c.JSON(200, gin.H{
		"id": id,
	})

}
