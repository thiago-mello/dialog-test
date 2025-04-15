package posts

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/middlewares"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/createpost"
)

func ConfigPostRoutes(router *echo.Echo) {
	// protected route group
	routeGroup := router.Group("/v1/posts", middlewares.RequireJWTAuth())

	routeGroup.POST("", createPost().Handle)
}

func createPost() *createpost.CreatePostHttpAdapter {
	return createpost.NewCreatePostAdapter(config.GetDb())
}
