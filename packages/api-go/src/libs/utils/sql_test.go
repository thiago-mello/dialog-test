package utils_test

import (
	"testing"

	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/stretchr/testify/assert"
)

func TestTranslateNamedQuery(t *testing.T) {
	query := "SELECT * FROM users WHERE name = :name AND age = :age"
	params := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	sql, args, err := utils.TranslateNamedQuery(query, params)
	assert.NoError(t, err)
	assert.Contains(t, sql, "$1")
	assert.Contains(t, sql, "$2")
	assert.Len(t, args, 2)
}
