package handler

import (
	"net/http"
	"github.com/agusbasari29/simru-be/p2ak3/helper"
	"github.com/agusbasari29/simru-be/p2ak3/request"
	"github.com/agusbasari29/simru-be/p2ak3/response"
	"github.com/agusbasari29/simru-be/p2ak3/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) Login(ctx *gin.Context) {
	var request request.AuthLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.authService.Login(request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	generateToken := service.NewJWTService().GenerateToken(user)
	userData := response.UserResponseFormatter(user)
	data := response.UserDataResponseFormatter(userData, generateToken)
	response := helper.ResponseFormatter(http.StatusOK, "success", nil, data)
	ctx.JSON(http.StatusOK, response)
}
