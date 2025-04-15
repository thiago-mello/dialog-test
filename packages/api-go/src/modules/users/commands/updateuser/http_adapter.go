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
