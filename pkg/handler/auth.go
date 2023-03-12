package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/todo-app"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input todo.User
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {

}
