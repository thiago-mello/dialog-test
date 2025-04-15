package createuser

type CreateUserCommand struct {
	Name     string
	Email    string
	Password string
	Bio      *string
}
