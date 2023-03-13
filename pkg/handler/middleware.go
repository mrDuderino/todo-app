package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationToken = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationToken)
	if header == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
