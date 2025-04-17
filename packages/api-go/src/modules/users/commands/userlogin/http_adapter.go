package userlogin

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/dto"
)

type UserLoginHttpAdapter struct {
	useCase UserLoginUseCase
}

func NewUserLoginHttpAdapter(db *sqlx.DB) *UserLoginHttpAdapter {
	return &UserLoginHttpAdapter{useCase: NewLoginUseCase(db)}
}

// Handle UserLogin godoc
//
//	@Summary		Authenticate user
//	@Description	Logs in a user using email and password
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	dto.LoginRequestDto	true	"Login credentials"
//	@Success		200		{object}	dto.LoginResponseDto
//	@Failure		400		{object}	errs.ErrorResponse
//	@Failure		401		{object}	errs.ErrorResponse
//	@Failure		500		{object}	errs.ErrorResponse
//	@Router			/v1/users/login [post]
func (c *UserLoginHttpAdapter) Handle(ctx echo.Context) error {
	body := dto.LoginRequestDto{}

	if err := utils.BindAndValidate(ctx, &body); err != nil {
		return err
	}

	command := UserLoginCommand{
		Email:    body.Email,
		Password: body.Password,
	}

	response, err := c.useCase.LoginUser(ctx.Request().Context(), command)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}
