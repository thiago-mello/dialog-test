package dto

type CreateUserDto struct {
	Name            string  `json:"name,omitempty" validate:"required"`
	Bio             *string `json:"bio,omitempty" validate:"omitempty,max=400"`
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password,omitempty" validate:"required,min=8"`
	PasswordConfirm string  `json:"password_confirm,omitempty" validate:"required,min=8,eqfield=Password"`
}
