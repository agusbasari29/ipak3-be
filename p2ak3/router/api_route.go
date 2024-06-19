package router

import (
	"github.com/agusbasari29/simru-be/p2ak3/helper"
	"github.com/agusbasari29/simru-be/p2ak3/middleware"
	"github.com/agusbasari29/simru-be/p2ak3/service"

	"github.com/gin-gonic/gin"
)

func DefineAuthRoutes(e *gin.Engine) {
	handlers := []helper.Handler{
		&AuthRoute{},
	}

	var routes []helper.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/auth")
	for _, route := range routes {
		api.POST(route.Path, route.HandlerFunc...)
	}
}

func DefineSecureRoutes(e *gin.Engine) {
	jwtService := service.NewJWTService()
	handlers := []helper.Handler{
		&UserRoute{},
	}

	var routes []helper.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api", middleware.AuthorizeJWT(jwtService))
	for _, route := range routes {
		switch route.Method {
		case "GET":
			api.GET(route.Path, route.HandlerFunc...)
		case "POST":
			api.POST(route.Path, route.HandlerFunc...)
		case "PUT":
			api.PUT(route.Path, route.HandlerFunc...)
		case "PATCH":
			api.PATCH(route.Path, route.HandlerFunc...)
		case "DELETE":
			api.DELETE(route.Path, route.HandlerFunc...)
		}
	}

}
