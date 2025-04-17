package utils_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePageSize(t *testing.T) {
	assert.Equal(t, int32(10), utils.CalculatePageSize(0))
	assert.Equal(t, int32(25), utils.CalculatePageSize(25))
}

func TestStringPointerToUuid(t *testing.T) {
	id := uuid.New().String()
	result := utils.StringPointerToUuid(&id)
	assert.NotNil(t, result)

	invalid := "not-a-uuid"
	assert.Nil(t, utils.StringPointerToUuid(&invalid))
	assert.Nil(t, utils.StringPointerToUuid(nil))
}
