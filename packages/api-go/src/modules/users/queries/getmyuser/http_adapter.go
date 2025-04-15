package getmyuser

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/dto"
)

type GetMyUserHttpAdapter struct {
	persistence database.UsersDatabaseOutputPort
}

func NewGetMyUserAdapter(db *sqlx.DB) *GetMyUserHttpAdapter {
	return &GetMyUserHttpAdapter{persistence: database.NewUsersDatabaseOutputPort(db)}
}

func (e *GetMyUserHttpAdapter) Query(ctx echo.Context) error {
	applicationContext := ctx.(*context.ApplicationContext)
	email := applicationContext.User.Email

	user, err := e.persistence.FindByEmail(ctx.Request().Context(), email)
	if err != nil {
		return err
	}
	if user == nil {
		return errs.BadRequestError("user was not found")
	}

	mappedUser := dto.GetUserResponseDto{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Bio:   user.Bio,
	}
	return ctx.JSON(http.StatusOK, mappedUser)
}
