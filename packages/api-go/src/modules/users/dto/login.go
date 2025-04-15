package dto

import "github.com/google/uuid"

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDto struct {
	AccessToken string          `json:"access_token"`
	User        UserResponseDto `json:"user"`
}

type UserResponseDto struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
