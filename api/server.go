package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joaofilippe/americanas-desafio/api/requests"
)

// Server is a simple handler
func Server() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	listGroup := api.Group("/v1")
	{
		listGroup.GET("/", HelloWorld)
		listGroup.POST("/save_list", SaveLists)
	}

	return r
}

// HelloWorld is a simple handler
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func SaveLists(c *gin.Context) {
	l := new(requests.SaveListsRequest)

	c.Bind(l)

	c.JSON(200, gin.H{
		"message": "it works",
		"data":    l,
	})
}
