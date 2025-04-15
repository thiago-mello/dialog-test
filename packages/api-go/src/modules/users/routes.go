package users

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/createuser"
)

func ConfigUserRoutes(router *echo.Echo) {
	routeGroup := router.Group("/v1/users")

	routeGroup.POST("", createUser().HandleCreateUserRequest)
}

func createUser() *createuser.CreateUserHttpAdapter {
	return createuser.NewCreateUserAdapter(config.GetDb())
}
