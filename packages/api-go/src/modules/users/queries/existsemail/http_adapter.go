package existsemail

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
)

type ExistsUserByEmailHttpAdapter struct {
	persistence database.UsersDatabaseOutputPort
}

func NewExistsUserByEmail(db *sqlx.DB) *ExistsUserByEmailHttpAdapter {
	return &ExistsUserByEmailHttpAdapter{persistence: database.NewUsersDatabaseOutputPort(db)}
}

// Query CheckEmailExists godoc
//
//	@Summary		Check if email exists
//	@Description	Returns whether a user with the given email exists
//	@Tags			Users
//	@Produce		json
//	@Param			email	query	string	true	"Email to check"
//	@Success		200		{object}	map[string]bool
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users/exists [get]
func (e *ExistsUserByEmailHttpAdapter) Query(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return errs.BadRequestError("email is required")
	}

	user, err := e.persistence.FindByEmail(ctx.Request().Context(), email)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]bool{"exists": user != nil})
}
