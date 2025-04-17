package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/createpost"
	"github.com/leandro-andrade-candido/api-go/src/tests/server"
	"github.com/leandro-andrade-candido/api-go/src/tests/testdb"
)

func TestCreatePost(t *testing.T) {
	t.Skip() // skip test
	db, err := testdb.SetupTestDatabase()
	assert.NoError(t, err)

	user := testdb.User

	e := server.GetServer()
	cacheMock := cache.NewRedisCache("localhost:6379", "", 0)
	adapter := createpost.NewCreatePostAdapter(db, cacheMock)

	req := httptest.NewRequest(http.MethodPost, "/v1/posts", strings.NewReader(`{"content": "<p>hello world</p>", "is_public": true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/posts")

	authCtx := &context.ApplicationContext{
		Context: c,
		User:    context.UserClaims{Id: user.ID, Email: user.Email},
	}

	err = adapter.Handle(authCtx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var res map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(t, "<p>hello world</p>", res["content"])
}
