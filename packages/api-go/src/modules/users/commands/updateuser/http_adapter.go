package updateuser

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/dto"
)

type UpdateCurrentUserHttpAdapter struct {
	useCase UpdateUserUseCase
}

func NewUpdateUserAdapter(db *sqlx.DB) *UpdateCurrentUserHttpAdapter {
	return &UpdateCurrentUserHttpAdapter{useCase: NewUpdateUserUseCase(db)}
}

// Handle UpdateUser godoc
//
//	@Summary		Update current user
//	@Description	Update the authenticated user with name, email, password, and bio
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	dto.UpdateUserDto	true	"User update payload"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		401		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users/me [put]
func (c *UpdateCurrentUserHttpAdapter) Handle(ctx echo.Context) error {
	body := &dto.UpdateUserDto{}
	applicationContext := ctx.(*context.ApplicationContext)

	err := utils.BindAndValidate(ctx, body)
	if err != nil {
		return err
	}

	command := UpdateUserCommand{
		UserId:   applicationContext.User.Id,
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Bio:      body.Bio,
	}

	err = c.useCase.UpdateUser(ctx.Request().Context(), command)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"id": command.UserId.String()})
}
