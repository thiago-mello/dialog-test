package posts

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/middlewares"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/createpost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/deletepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/likepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/unlikepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/updatepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/queries/getpost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/queries/listmyposts"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/queries/listposts"
)

func ConfigPostRoutes(router *echo.Echo) {
	// protected route group
	routeGroup := router.Group("/v1/posts", middlewares.RequireJWTAuth())

	//posts
	routeGroup.POST("", createPost().Handle)
	routeGroup.PUT("/:id", updatePost().Handle)
	routeGroup.GET("/:id", getPost().Query)
	routeGroup.GET("", listPosts().Query)
	routeGroup.GET("/my-posts", listMyPosts().Query)
	routeGroup.DELETE("/:id", deletePost().Handle)

	// likes
	routeGroup.POST("/:id/likes", likePost().Handle)
	routeGroup.DELETE("/:id/likes", unlikePost().Handle)

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

func listPosts() *listposts.ListPostsHttpAdapter {
	return listposts.NewListPostAdapter(config.GetDb())
}

func listMyPosts() *listmyposts.ListMyPostsHttpAdapter {
	return listmyposts.NewListMyPostAdapter(config.GetDb())
}

func deletePost() *deletepost.DeletePostHttpAdapter {
	return deletepost.NewDeletePostAdapter(config.GetDb())
}

func likePost() *likepost.LikePostHttpAdapter {
	return likepost.NewLikePostAdapter(config.GetDb())
}

func unlikePost() *unlikepost.UnlikePostHttpAdapter {
	return unlikepost.NewUnlikePostAdapter(config.GetDb())
}
