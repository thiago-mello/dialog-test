package utils

import "github.com/labstack/echo/v4"

// BindAndValidate binds request data to the provided destination struct and validates it
// Parameters:
//   - ctx: Echo context containing the request data
//   - dest: Pointer to destination struct where request data will be bound
//
// Returns:
//   - error: Returns nil if binding and validation succeed, otherwise returns the error
func BindAndValidate(ctx echo.Context, dest any) error {
	if err := ctx.Bind(dest); err != nil {
		return err
	}
	if err := ctx.Validate(dest); err != nil {
		return err
	}

	return nil
}
