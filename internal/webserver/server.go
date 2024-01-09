package webserver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaofilippe/americanas-desafio/internal/common"
	"github.com/joaofilippe/americanas-desafio/internal/interfaces"
	"github.com/joaofilippe/americanas-desafio/internal/models"
	"github.com/joaofilippe/americanas-desafio/internal/webserver/request"
	"github.com/joaofilippe/americanas-desafio/internal/webserver/response"
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
	l := new(request.SaveListsRequest)

	err := c.Bind(l)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: http.StatusBadRequest,
			ErrorMessage:  common.ErrMsgContentNotBinded,
			Data: err,
		})
		return
	}

	id, err := w.Application.SaveListsNode(l.List1, l.List2)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code: http.StatusInternalServerError,
			ErrorMessage: common.ErrMsgCannotSaveLists,
			Data: err,
		})
		return
	}

	c.JSON(http.StatusCreated, 
		response.Response{
			Code: http.StatusCreated,
			Message: "Lists saved successfully. Follow the id to merge the lists.",
			Data: id,
		})
}

// MergeLists is a simple handler
func (w *WebApp) MergeLists(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	list, err := w.Application.MergeListNode(id)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code: http.StatusInternalServerError,
			ErrorMessage: common.ErrMsgCannotMergeLists,
			Data: err,
		})
		return
	}


	c.JSON(http.StatusCreated, 
		response.Response{
			Code: http.StatusCreated,
			Message: "Lists had merged succesfully.",
			Data: struct{
				MergedList *models.ListNode `json:"merged_list"`
			}{
				MergedList: list,
			},
		})

}
