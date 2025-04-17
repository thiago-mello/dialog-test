package utils_test

import (
	"testing"

	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestIsConstraintViolation(t *testing.T) {
	err := &pq.Error{Constraint: "unique_email"}
	assert.True(t, utils.IsConstraintViolation(err, "unique_email"))
	assert.False(t, utils.IsConstraintViolation(err, "other_constraint"))
	assert.False(t, utils.IsConstraintViolation(nil, "unique_email"))
}
