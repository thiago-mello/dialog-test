package utils_test

import (
	"testing"

	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "secret123"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	match := utils.VerifyPasswordHash(password, hash)
	assert.True(t, match)

	fail := utils.VerifyPasswordHash("wrong", hash)
	assert.False(t, fail)
}
