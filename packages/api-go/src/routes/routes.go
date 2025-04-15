package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts"
	"github.com/leandro-andrade-candido/api-go/src/modules/users"
)

// SetupRoutes configures all API routes by calling the route configuration
// functions from each module
// Parameters:
//   - e: Echo instance used to register the routes
func SetupRoutes(e *echo.Echo) {
	users.ConfigUserRoutes(e)
	posts.ConfigPostRoutes(e)
}
