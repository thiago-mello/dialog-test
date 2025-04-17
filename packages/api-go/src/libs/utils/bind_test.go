package utils_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name string `json:"name" validate:"required"`
}

type fakeValidator struct{}

func (f *fakeValidator) Validate(i any) error {
	if v, ok := i.(*testStruct); ok && v.Name == "" {
		return errors.New("validation failed")
	}
	return nil
}

func TestBindAndValidate_Success(t *testing.T) {
	e := echo.New()
	e.Validator = &fakeValidator{}
	body := `{"name":"Echo"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var data testStruct
	err := utils.BindAndValidate(c, &data)

	assert.NoError(t, err)
	assert.Equal(t, "Echo", data.Name)
}

func TestBindAndValidate_ValidationError(t *testing.T) {
	e := echo.New()
	e.Validator = &fakeValidator{}
	body := `{"name":""}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var data testStruct
	err := utils.BindAndValidate(c, &data)

	assert.Error(t, err)
}
