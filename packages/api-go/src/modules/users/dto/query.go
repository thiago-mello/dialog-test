package dto

import "github.com/google/uuid"

type GetUserResponseDto struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
	Bio   *string   `json:"bio,omitempty"`
}
