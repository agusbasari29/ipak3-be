package router

import (
	"github.com/agusbasari29/simru-be/p2ak3/database"
	"github.com/agusbasari29/simru-be/p2ak3/handler"
	"github.com/agusbasari29/simru-be/p2ak3/helper"
	"github.com/agusbasari29/simru-be/p2ak3/repository"
	"github.com/agusbasari29/simru-be/p2ak3/service"

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
