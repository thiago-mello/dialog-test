package deleteuser

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
)

type DeleteUserHttpAdapter struct {
	useCase DeleteUserUseCase
}

func NewDeleteUserAdapter(db *sqlx.DB, cache cache.Cache) *DeleteUserHttpAdapter {
	return &DeleteUserHttpAdapter{useCase: NewDeleteUserUseCase(db, cache)}
}

// Handle DeleteUser godoc
//
//	@Summary		Removes an user account
//	@Description	Removes an user account and its associated data
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Success		200
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users/me [delete]
func (a *DeleteUserHttpAdapter) Handle(ctx echo.Context) error {
	applicationContext := ctx.(*context.ApplicationContext)
	command := DeleteUserCommand{
		UserId: applicationContext.User.Id,
	}

	err := a.useCase.DeleteUser(ctx.Request().Context(), command)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
