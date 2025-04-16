package posts

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/middlewares"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/createpost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/updatepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/queries/getpost"
)

func ConfigPostRoutes(router *echo.Echo) {
	// protected route group
	routeGroup := router.Group("/v1/posts", middlewares.RequireJWTAuth())

	routeGroup.POST("", createPost().Handle)
	routeGroup.PUT("/:id", updatePost().Handle)
	routeGroup.GET("/:id", getPost().Query)
}

func createPost() *createpost.CreatePostHttpAdapter {
	return createpost.NewCreatePostAdapter(config.GetDb())
}

func updatePost() *updatepost.UpdatePostHttpAdapter {
	return updatepost.NewUpdatePostAdapter(config.GetDb())
}

func getPost() *getpost.GetPostHttpAdapter {
	return getpost.NewGetPostAdapter(config.GetDb())
}
