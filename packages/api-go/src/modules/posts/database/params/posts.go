package params

import "github.com/google/uuid"

type GetPostsParams struct {
	PageSize      int32      `db:"page_size"`
	LastSeenId    *uuid.UUID `db:"last_seen_id"`
	CurrentUserId uuid.UUID  `db:"current_user_id"`
	UserId        *uuid.UUID `db:"user_id"`
}
