package createuser

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/dto"
)

type CreateUserHttpAdapter struct {
	useCase CreateUserUseCase
}

func NewCreateUserAdapter(db *sqlx.DB) *CreateUserHttpAdapter {
	return &CreateUserHttpAdapter{useCase: NewUseCase(db)}
}

// Handle CreateUser godoc
//
//	@Summary		Register a new user
//	@Description	Creates a user with name, email, password and optional bio
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	dto.CreateUserDto	true	"User payload"
//	@Success		201
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users [post]
func (c *CreateUserHttpAdapter) Handle(ctx echo.Context) error {
	body := &dto.CreateUserDto{}

	err := ctx.Bind(&body)
	if err != nil {
		return err
	}

	err = ctx.Validate(body)
	if err != nil {
		return err
	}

	command := CreateUserCommand{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Bio:      body.Bio,
	}

	err = c.useCase.CreateNewUser(ctx.Request().Context(), command)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
