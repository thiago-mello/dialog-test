package dto

type CreatePostDto struct {
	Content  string `json:"content" validate:"required,max=12000"`
	IsPublic bool   `json:"is_public" validate:"required"`
}

type PostCreatedResponseDto struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	IsPublic bool   `json:"is_public"`
}
