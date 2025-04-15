package users

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/createuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/userlogin"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/queries/existsemail"
)

func ConfigUserRoutes(router *echo.Echo) {
	routeGroup := router.Group("/v1/users")

	routeGroup.POST("", createUser().Handle)
	routeGroup.GET("/exists", existsUser().Query)
	routeGroup.POST("/login", login().Handle)
}

func createUser() *createuser.CreateUserHttpAdapter {
	return createuser.NewCreateUserAdapter(config.GetDb())
}

func existsUser() *existsemail.ExistsUserByEmailHttpAdapter {
	return existsemail.NewExistsUserByEmail(config.GetDb())
}

func login() *userlogin.UserLoginHttpAdapter {
	return userlogin.NewUserLoginHttpAdapter(config.GetDb())
}
