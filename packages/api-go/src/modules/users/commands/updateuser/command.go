package updateuser

import "github.com/google/uuid"

type UpdateUserCommand struct {
	UserId   uuid.UUID
	Name     string
	Email    string
	Bio      *string
	Password string
}
