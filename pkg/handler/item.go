package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/todo-app/models"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input models.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.CreateItem(userId, listId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type getAllItemsResponse struct {
	Data []models.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllItemsResponse{Data: items})
}

func (h *Handler) getItemById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(ctx *gin.Context) {

}

func (h *Handler) deleteItem(ctx *gin.Context) {

}
