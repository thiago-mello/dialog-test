package updatepost_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/leandro-andrade-candido/api-go/src/modules/posts/commands/updatepost"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/params"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/projections"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

// MockPostsDatabase implements the PostsDatabaseOutputPort interface
type MockPostsDatabase struct {
	mock.Mock
}

func (m *MockPostsDatabase) Save(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	return m.Called(ctx, tx, post).Error(0)
}

func (m *MockPostsDatabase) FindByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	args := m.Called(ctx, id)
	var post any = args.Get(0)
	if post == nil {
		return nil, args.Error(1)
	}

	return post.(*domain.Post), args.Error(1)
}

func (m *MockPostsDatabase) Update(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	return m.Called(ctx, tx, post).Error(0)
}

func (m *MockPostsDatabase) ListPosts(ctx context.Context, filters params.GetPostsParams) (*[]projections.ListPostsProjection, error) {
	return nil, nil
}

func (m *MockPostsDatabase) Delete(ctx context.Context, tx *sqlx.Tx, postID, userID uuid.UUID) error {
	return nil
}

func TestUpdatePostService_UpdatePost_Success(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	userID := uuid.New()
	postID := uuid.New()
	existingPost := &domain.Post{
		ID:       postID,
		UserID:   userID,
		Content:  "Original content",
		IsPublic: false,
	}

	command := updatepost.UpdatePostCommand{
		PostID:   postID,
		UserID:   userID,
		Content:  "Updated content",
		IsPublic: true,
	}

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, postID).
		Return(existingPost, nil).
		Once()

	mockDB.On("Update", mock.Anything, (*sqlx.Tx)(nil), mock.AnythingOfType("*domain.Post")).
		Run(func(args mock.Arguments) {
			updatedPost := args.Get(2).(*domain.Post)
			assert.Equal(t, command.Content, updatedPost.Content)
			assert.Equal(t, command.IsPublic, updatedPost.IsPublic)
		}).
		Return(nil).
		Once()

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, command.Content, result.Content)
	assert.Equal(t, command.IsPublic, result.IsPublic)
	mockDB.AssertExpectations(t)
}

func TestUpdatePostService_UpdatePost_PostNotFound(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	command := updatepost.UpdatePostCommand{
		PostID: uuid.New(),
		UserID: uuid.New(),
	}

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, command.PostID).
		Return(nil, nil).
		Once()

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "post not found")
	mockDB.AssertExpectations(t)
}

func TestUpdatePostService_UpdatePost_UnauthorizedUser(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	postID := uuid.New()
	existingPost := &domain.Post{
		ID:     postID,
		UserID: uuid.New(), // Different user ID
	}

	command := updatepost.UpdatePostCommand{
		PostID: postID,
		UserID: uuid.New(), // Requesting user ID
	}

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, postID).
		Return(existingPost, nil).
		Once()

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "post not found")
	mockDB.AssertExpectations(t)
}

func TestUpdatePostService_UpdatePost_ValidationFailure(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	userID := uuid.New()
	postID := uuid.New()
	existingPost := &domain.Post{
		ID:     postID,
		UserID: userID,
	}

	command := updatepost.UpdatePostCommand{
		PostID:  postID,
		UserID:  userID,
		Content: "", // Invalid empty content
	}

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, postID).
		Return(existingPost, nil).
		Once()

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "content cannot be empty")
	mockDB.AssertExpectations(t)
}

func TestUpdatePostService_UpdatePost_DatabaseErrorOnFind(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	postID := uuid.New()
	expectedErr := errors.New("database connection failed")

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, postID).
		Return(&domain.Post{}, expectedErr).
		Once()

	command := updatepost.UpdatePostCommand{
		PostID: postID,
		UserID: uuid.New(),
	}

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}

func TestUpdatePostService_UpdatePost_DatabaseErrorOnUpdate(t *testing.T) {
	// Setup
	mockDB := new(MockPostsDatabase)
	service := updatepost.UpdatePostService{
		Persistence: mockDB,
	}

	// Test data
	userID := uuid.New()
	postID := uuid.New()
	existingPost := &domain.Post{
		ID:     postID,
		UserID: userID,
	}
	expectedErr := errors.New("update failed")

	command := updatepost.UpdatePostCommand{
		PostID:  postID,
		UserID:  userID,
		Content: "Valid content",
	}

	// Mock expectations
	mockDB.On("FindByID", mock.Anything, postID).
		Return(existingPost, nil).
		Once()

	mockDB.On("Update", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedErr).
		Once()

	// Execution
	result, err := service.UpdatePost(context.Background(), command)

	// Assertions
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}
