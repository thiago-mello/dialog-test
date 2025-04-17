package createuser_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/createuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
)

type MockUsersDatabase struct {
	mock.Mock
}

func (m *MockUsersDatabase) Insert(ctx context.Context, tx *sqlx.Tx, user *domain.User) (bool, error) {
	args := m.Called(ctx, tx, user)
	return args.Bool(0), args.Error(1)
}

func (m *MockUsersDatabase) UpdateById(ctx context.Context, tx *sqlx.Tx, user *domain.User) (bool, error) {
	args := m.Called(ctx, tx, user)
	return args.Bool(0), args.Error(1)
}

func (m *MockUsersDatabase) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	var user any = args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}

	return user.(*domain.User), args.Error(1)
}

func TestCreateUserService_CreateNewUser_Success(t *testing.T) {
	// Setup
	mockDB := new(MockUsersDatabase)
	service := createuser.CreateUserService{
		Persistence: mockDB,
	}

	// Test data
	validCommand := createuser.CreateUserCommand{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securePassword123!",
		Bio:      lo.ToPtr("Software Developer"),
	}

	// Mock expectations
	mockDB.On("Insert", mock.Anything, mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
		return user.Email == validCommand.Email &&
			user.Name == validCommand.Name &&
			*user.Bio == *validCommand.Bio &&
			user.PasswordHash != ""
	})).
		Return(false, nil). // userExists = false
		Once()

	// Execution
	err := service.CreateNewUser(context.Background(), validCommand)

	// Assertions
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCreateUserService_CreateNewUser_EmailConflict(t *testing.T) {
	// Setup
	mockDB := new(MockUsersDatabase)
	service := createuser.CreateUserService{
		Persistence: mockDB,
	}

	// Test data
	validCommand := createuser.CreateUserCommand{
		Name:     "John Doe",
		Email:    "existing@example.com",
		Password: "securePassword123!",
	}

	// Mock expectations
	mockDB.On("Insert", mock.Anything, mock.Anything, mock.Anything).
		Return(true, nil).
		Once()

	// Execution
	err := service.CreateNewUser(context.Background(), validCommand)

	// Assertions
	assert.Error(t, err)
	assert.IsType(t, errs.ApiError{}, err)
	assert.Equal(t, "The email requested is not available", err.Error())
	mockDB.AssertExpectations(t)
}

func TestCreateUserService_CreateNewUser_ValidationFailure(t *testing.T) {
	testCases := []struct {
		name        string
		command     createuser.CreateUserCommand
		expectedErr string
	}{
		{
			name: "Empty Name",
			command: createuser.CreateUserCommand{
				Name:     "",
				Email:    "john@example.com",
				Password: "password",
			},
			expectedErr: "name is mandatory",
		},
		{
			name: "Invalid Email",
			command: createuser.CreateUserCommand{
				Name:     "John",
				Email:    "invalid-email",
				Password: "password",
			},
			expectedErr: "invalid email",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			service := createuser.CreateUserService{}
			err := service.CreateNewUser(context.Background(), tc.command)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}

func TestCreateUserService_CreateNewUser_DatabaseError(t *testing.T) {
	// Setup
	mockDB := new(MockUsersDatabase)
	service := createuser.CreateUserService{
		Persistence: mockDB,
	}

	// Test data
	validCommand := createuser.CreateUserCommand{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securePassword123!",
	}

	// Mock expectations
	expectedErr := errors.New("database connection failed")
	mockDB.On("Insert", mock.Anything, mock.Anything, mock.Anything).
		Return(false, expectedErr).
		Once()

	// Execution
	err := service.CreateNewUser(context.Background(), validCommand)

	// Assertions
	assert.ErrorIs(t, err, expectedErr)
	mockDB.AssertExpectations(t)
}
