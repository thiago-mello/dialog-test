package context

import "github.com/labstack/echo/v4"

type ApplicationContext struct {
	echo.Context
	User UserClaims
}

func (c ApplicationContext) IsUserAuthenticated() bool {
	return c.User.Email != ""
}
