package userlogin

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/config/token"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/dto"
)

type UserLoginUseCase interface {
	// Creates and persists new user into the database
	LoginUser(ctx context.Context, command UserLoginCommand) (*dto.LoginResponseDto, error)
}

type UserLoginService struct {
	persistence database.UsersDatabaseOutputPort
}

func NewLoginUseCase(db *sqlx.DB) UserLoginUseCase {
	return &UserLoginService{
		persistence: database.NewUsersDatabaseOutputPort(db),
	}
}

func (u *UserLoginService) LoginUser(ctx context.Context, command UserLoginCommand) (*dto.LoginResponseDto, error) {
	lowerCaseEmail := strings.ToLower(command.Email)
	user, err := u.persistence.FindByEmail(ctx, lowerCaseEmail)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NotFoundError("Email or password is incorrect")
	}

	validPassword := utils.VerifyPasswordHash(command.Password, user.PasswordHash)
	if !validPassword {
		return nil, errs.NotFoundError("Email or password is incorrect")
	}

	tokenString, err := signTokenForUser(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		AccessToken: tokenString,
		User: dto.UserResponseDto{
			Id:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

// Generates a JWT token for the given user
// Parameters:
//   - user: A pointer to a domain.User struct containing user information
//
// Returns:
//   - string: The signed JWT token string (HS256 algorithm)
//   - error: An error if token generation fails, nil on success
func signTokenForUser(user *domain.User) (string, error) {

	tokenConfig, err := token.GetTokenConfiguration()
	if err != nil {
		return "", err
	}

	tokenClaims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(tokenConfig.ExpiresIn).Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.New().String(),

		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString([]byte(tokenConfig.Secret))
}
