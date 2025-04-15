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

func (e *ExistsUserByEmailHttpAdapter) Query(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return errs.BadRequestError("email is required")
	}

	exists, err := e.persistence.ExistsByEmail(ctx.Request().Context(), email)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]bool{"exists": exists})
}
