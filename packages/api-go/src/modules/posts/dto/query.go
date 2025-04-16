package dto

import (
	"time"
)

type PostResponseDto struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ListPostRequestDto struct {
	PageSize   int     `query:"page_size" validate:"omitempty,min=1,max=50"`
	LastSeenId *string `query:"last_seen_id"`
}

type ListPostResponseDto struct {
	Id                string          `json:"id,omitempty"`
	Content           string          `json:"content,omitempty"`
	IsPrivate         bool            `json:"is_private"`
	User              ListPostUserDto `json:"user"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	LikeCount         int32           `json:"like_count"`
	UserLikedThisPost bool            `json:"user_liked_this_post"`
}

type ListPostUserDto struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}
