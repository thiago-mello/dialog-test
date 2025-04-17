package middlewares

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/config/token"
	"github.com/stretchr/testify/assert"
)

func TestParseJwtToken(t *testing.T) {
	secret := "mysecret"
	userID := uuid.New()
	email := "test@example.com"

	validToken := func() string {
		claims := jwt.MapClaims{
			"sub":   userID.String(),
			"email": email,
			"exp":   time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString([]byte(secret))
		return signedToken
	}()

	tests := []struct {
		name           string
		token          string
		configProvider ConfigProvider
		expectErr      bool
		expectedUserID uuid.UUID
		expectedEmail  string
	}{
		{
			name:  "valid token",
			token: validToken,
			configProvider: func() (*token.TokenConfig, error) {
				return &token.TokenConfig{Secret: secret}, nil
			},
			expectErr:      false,
			expectedUserID: userID,
			expectedEmail:  email,
		},
		{
			name: "invalid signing method",
			token: func() string {
				claims := jwt.MapClaims{"sub": userID.String(), "email": email}
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				signedToken, _ := token.SignedString([]byte(secret)) // Incorrect but sufficient for test
				return signedToken
			}(),
			configProvider: func() (*token.TokenConfig, error) {
				return &token.TokenConfig{Secret: secret}, nil
			},
			expectErr: true,
		},
		{
			name:  "error from config provider",
			token: validToken,
			configProvider: func() (*token.TokenConfig, error) {
				return &token.TokenConfig{}, errors.New("config error")
			},
			expectErr: true,
		},
		{
			name: "invalid UUID in sub",
			token: func() string {
				claims := jwt.MapClaims{"sub": "not-a-uuid", "email": email}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				signedToken, _ := token.SignedString([]byte(secret))
				return signedToken
			}(),
			configProvider: func() (*token.TokenConfig, error) {
				return &token.TokenConfig{Secret: secret}, nil
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			claims, err := ParseJwtToken(tc.token, tc.configProvider)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedUserID, claims.Id)
				assert.Equal(t, tc.expectedEmail, claims.Email)
			}
		})
	}
}
