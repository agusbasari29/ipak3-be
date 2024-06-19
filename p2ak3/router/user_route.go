package router

import (
	"github.com/agusbasari29/simru-be/p2ak3/database"
	"github.com/agusbasari29/simru-be/p2ak3/handler"
	"github.com/agusbasari29/simru-be/p2ak3/helper"
	"github.com/agusbasari29/simru-be/p2ak3/repository"
	"github.com/agusbasari29/simru-be/p2ak3/service"

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
