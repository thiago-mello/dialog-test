package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		hasError bool
	}{
		{
			name: "valid user",
			user: User{
				ID:           uuid.New(),
				Email:        "user@example.com",
				PasswordHash: "somehash",
				Name:         "John Doe",
			},
			hasError: false,
		},
		{
			name: "missing email",
			user: User{
				ID:           uuid.New(),
				Email:        "",
				PasswordHash: "somehash",
				Name:         "John Doe",
			},
			hasError: true,
		},
		{
			name: "invalid email format",
			user: User{
				ID:           uuid.New(),
				Email:        "invalid-email",
				PasswordHash: "somehash",
				Name:         "John Doe",
			},
			hasError: true,
		},
		{
			name: "missing password hash",
			user: User{
				ID:           uuid.New(),
				Email:        "user@example.com",
				PasswordHash: "",
				Name:         "John Doe",
			},
			hasError: true,
		},
		{
			name: "missing name",
			user: User{
				ID:           uuid.New(),
				Email:        "user@example.com",
				PasswordHash: "somehash",
				Name:         "",
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
