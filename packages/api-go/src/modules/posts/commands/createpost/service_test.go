package createpost_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/createpost"
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

func TestCreatePostService_CreatePost_Success(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := createpost.CreatePostService{
		Persistence: mockDB,
	}

	// Test data
	userID := uuid.New()
	validCommand := createpost.CreatePostCommand{
		UserID:   userID,
		Content:  "Valid content",
		IsPublic: true,
	}

	// Mock expectations
	mockDB.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("*domain.Post")).
		Return(nil).
		Once()

	// Execution
	result, err := service.CreatePost(context.Background(), validCommand)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, validCommand.Content, result.Content)
	mockDB.AssertExpectations(t)
}

func TestCreatePostService_CreatePost_ValidationFailure(t *testing.T) {
	// Setup
	service := createpost.CreatePostService{} // No DB dependency needed for validation test

	// Test data
	invalidCommand := createpost.CreatePostCommand{
		UserID:   uuid.New(),
		Content:  "", // Invalid empty content
		IsPublic: true,
	}

	// Execution
	result, err := service.CreatePost(context.Background(), invalidCommand)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "content cannot be empty")
}

func TestCreatePostService_Error_When_Saving(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := createpost.CreatePostService{
		Persistence: mockDB,
	}

	command := createpost.CreatePostCommand{
		UserID:   uuid.New(),
		Content:  "Valid content",
		IsPublic: true,
	}

	// Mock expectations
	mockDB.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("*domain.Post")).
		Return(errors.New("generated error")).
		Once()

		// Execution
	result, err := service.CreatePost(context.Background(), command)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "generated error")
}
