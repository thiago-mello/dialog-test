package context

type UserClaims struct {
	Email string `json:"email"`
	Id    uint32 `json:"id"`
}
