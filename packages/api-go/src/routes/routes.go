package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/modules/users"
)

func SetupRoutes(e *echo.Echo) {
	users.ConfigUserRoutes(e)
}
