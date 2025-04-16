package dto

import "time"

type PostResponseDto struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
