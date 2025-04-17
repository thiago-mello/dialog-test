package utils_test

import (
	"testing"

	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeHTML(t *testing.T) {
	input := `<script>alert("xss")</script><b>bold</b>`
	sanitized := utils.SanitizeHTML(input)

	assert.NotContains(t, sanitized, "script")
	assert.Contains(t, sanitized, "<b>bold</b>")
}
