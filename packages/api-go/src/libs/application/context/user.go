package context

import "github.com/google/uuid"

type UserClaims struct {
	Email string    `json:"email"`
	Id    uuid.UUID `json:"id"`
}
