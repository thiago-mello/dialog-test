package utils

import "github.com/labstack/echo/v4"

func BindAndValidate(ctx echo.Context, dest interface{}) error {
	if err := ctx.Bind(dest); err != nil {
		return err
	}
	if err := ctx.Validate(dest); err != nil {
		return err
	}

	return nil
}
