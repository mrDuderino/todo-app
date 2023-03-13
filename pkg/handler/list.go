package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/todo-app/models"
	"net/http"
)

func (h *Handler) createList(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input models.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateList(userId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getAllLists(ctx *gin.Context) {

}

func (h *Handler) getListById(ctx *gin.Context) {

}

func (h *Handler) updateList(ctx *gin.Context) {

}

func (h *Handler) deleteList(ctx *gin.Context) {

}
