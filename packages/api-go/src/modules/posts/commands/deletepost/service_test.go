package deletepost_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/deletepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/params"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/projections"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostsDatabase struct {
	mock.Mock
}

func (m *MockPostsDatabase) Save(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	args := m.Called(ctx, tx, post)
	return args.Error(0)
}

// Implement other interface methods with mock.Anything for this test
func (m *MockPostsDatabase) FindByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Post), args.Error(1)
}
func (m *MockPostsDatabase) Update(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	args := m.Called(ctx, tx, post)
	return args.Error(0)
}
func (m *MockPostsDatabase) ListPosts(ctx context.Context, filters params.GetPostsParams) (*[]projections.ListPostsProjection, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).(*[]projections.ListPostsProjection), args.Error(1)
}
func (m *MockPostsDatabase) Delete(ctx context.Context, tx *sqlx.Tx, postID, userID uuid.UUID) error {
	args := m.Called(ctx, tx, postID, userID)
	return args.Error(0)
}

func Test_DeletePostService_DeletePost_ValidDeletion(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := deletepost.DeletePostService{
		Persistence: mockDB,
	}

	// Expectations
	mockDB.On("Delete", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	command := deletepost.DeletePostCommand{
		PostID: uuid.New(),
		UserID: uuid.New(),
	}
	//Result
	err := service.DeletePost(context.Background(), command)

	assert.NoError(t, err)
}

func Test_DeletePostService_DeletePost_PostNotFound(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := deletepost.DeletePostService{
		Persistence: mockDB,
	}

	// Expectations
	mockDB.On("Delete", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("no posts deleted"))

	command := deletepost.DeletePostCommand{
		PostID: uuid.New(),
		UserID: uuid.New(),
	}
	//Result
	err := service.DeletePost(context.Background(), command)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Post not found")
}
