package users

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/middlewares"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/createuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/deleteuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/updateuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/userlogin"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/queries/existsemail"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/queries/getmyuser"
)

func ConfigUserRoutes(router *echo.Echo) {
	routeGroup := router.Group("/v1/users")

	routeGroup.POST("", createUser().Handle)
	routeGroup.GET("/exists", existsUser().Query)
	routeGroup.POST("/login", login().Handle)

	// protected user routes
	routeGroup.PUT("/me", updateUser().Handle, middlewares.RequireJWTAuth())
	routeGroup.GET("/me", getMyUser().Query, middlewares.RequireJWTAuth())
	routeGroup.DELETE("/me", deleteMyUser().Handle, middlewares.RequireJWTAuth())
}

func createUser() *createuser.CreateUserHttpAdapter {
	return createuser.NewCreateUserAdapter(config.GetDb())
}

func existsUser() *existsemail.ExistsUserByEmailHttpAdapter {
	return existsemail.NewExistsUserByEmail(config.GetDb())
}

func updateUser() *updateuser.UpdateCurrentUserHttpAdapter {
	return updateuser.NewUpdateUserAdapter(config.GetDb())
}

func getMyUser() *getmyuser.GetMyUserHttpAdapter {
	return getmyuser.NewGetMyUserAdapter(config.GetDb())
}

func login() *userlogin.UserLoginHttpAdapter {
	return userlogin.NewUserLoginHttpAdapter(config.GetDb())
}

func deleteMyUser() *deleteuser.DeleteUserHttpAdapter {
	return deleteuser.NewDeleteUserAdapter(config.GetDb(), config.GetCache())
}
