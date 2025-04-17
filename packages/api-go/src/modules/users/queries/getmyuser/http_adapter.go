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

// Query GetCurrentUser godoc
//
//	@Summary		Get current user
//	@Description	Fetch details of the authenticated user
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Success		200		{object}	dto.GetUserResponseDto
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		401		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users/me [get]
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
