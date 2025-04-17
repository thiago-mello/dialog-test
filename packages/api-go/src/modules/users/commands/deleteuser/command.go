package deleteuser

import "github.com/google/uuid"

type DeleteUserCommand struct {
	UserId uuid.UUID
}
