package router

import (
	"github.com/agusbasari29/simru-be/database"
	"github.com/agusbasari29/simru-be/handler"
	"github.com/agusbasari29/simru-be/helper"
	"github.com/agusbasari29/simru-be/repository"
	"github.com/agusbasari29/simru-be/service"

	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (r UserRoute) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepo := repository.NewUserRepository(db)
	contactRepo := repository.NewContactRepository(db)
	userService := service.NewUserService(userRepo, contactRepo)
	userHandler := handler.NewUserHandler(userService)
	return []helper.Route{
		{
			Method:      "GET",
			Path:        "/user/profile/:id",
			HandlerFunc: []gin.HandlerFunc{userHandler.GetByID},
		},
		{
			Method:      "PUT",
			Path:        "/user/update",
			HandlerFunc: []gin.HandlerFunc{userHandler.Create},
		},
	}
}
