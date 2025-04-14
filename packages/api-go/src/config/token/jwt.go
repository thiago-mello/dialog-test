package token

import (
	"errors"
	"strings"
	"time"

	"github.com/leandro-andrade-candido/api-go/src/config"
)

type TokenConfig struct {
	ExpiresIn time.Duration
	Secret    string
}

// Retrieves and validates JWT token configuration settings.
// It reads the secret key and expiration duration from the configuration.
//
// Returns:
//   - *TokenConfig: Contains the validated token secret and expiration duration
//   - error: Returns an error if the secret is empty or expiration duration is invalid
func GetTokenConfiguration() (*TokenConfig, error) {
	secret := config.GetString("zero.jwt.secret")
	if strings.Trim(secret, " ") == "" {
		return nil, errors.New("invalid secret for token")
	}

	duration, err := time.ParseDuration(config.GetString("zero.jwt.expires-in"))
	if err != nil {
		return nil, errors.New("invalid expiration date for token")
	}

	return &TokenConfig{
		ExpiresIn: duration,
		Secret:    secret,
	}, nil
}
