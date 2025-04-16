package dto

type UpdatePostDto struct {
	Content  string `json:"content" validate:"required,max=12000"`
	IsPublic bool   `json:"is_public"`
}

type PostUpdatedResponseDto struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	IsPublic bool   `json:"is_public"`
}
