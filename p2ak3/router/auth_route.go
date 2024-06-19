package router

import (
	"github.com/agusbasari29/simru-be/database"
	"github.com/agusbasari29/simru-be/handler"
	"github.com/agusbasari29/simru-be/helper"
	"github.com/agusbasari29/simru-be/repository"
	"github.com/agusbasari29/simru-be/service"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct{}

func (r AuthRoute) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	AuthHandler := handler.NewAuthHandler(authService)

	return []helper.Route{
		{
			Method:      "POST",
			Path:     "/login",
			HandlerFunc: []gin.HandlerFunc{AuthHandler.Login},
		},
	}
}
