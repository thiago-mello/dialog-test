package updateuser_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/commands/updateuser"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
)

// Mock for UsersDatabaseOutputPort
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
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUsersDatabase) DeleteById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (int64, error) {
	args := m.Called(ctx, tx, id)
	return args.Get(0).(int64), args.Error(1)
}

// Factory for UpdateUserService with mocked persistence
func newServiceWithMock(mockDB *MockUsersDatabase) updateuser.UpdateUserUseCase {
	return &updateuser.UpdateUserService{
		Persistence: mockDB,
	}
}

func TestUpdateUser_Success(t *testing.T) {
	mockDB := new(MockUsersDatabase)
	service := newServiceWithMock(mockDB)

	cmd := updateuser.UpdateUserCommand{
		UserId: uuid.New(),
		Email:  "new@example.com",
		Name:   "New Name",
		Bio:    utils.StringPointer("updated bio"),
	}

	mockDB.On("UpdateById", mock.Anything, (*sqlx.Tx)(nil), mock.AnythingOfType("*domain.User")).
		Return(false, nil)

	err := service.UpdateUser(context.Background(), cmd)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateUser_WithPassword(t *testing.T) {
	mockDB := new(MockUsersDatabase)
	service := newServiceWithMock(mockDB)

	cmd := updateuser.UpdateUserCommand{
		UserId:   uuid.New(),
		Email:    "withpass@example.com",
		Name:     "With Password",
		Bio:      utils.StringPointer("bio"),
		Password: "securepassword",
	}

	mockDB.On("UpdateById", mock.Anything, (*sqlx.Tx)(nil), mock.MatchedBy(func(user *domain.User) bool {
		return user.Email == cmd.Email && user.PasswordHash != ""
	})).Return(false, nil)

	err := service.UpdateUser(context.Background(), cmd)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateUser_EmailTaken(t *testing.T) {
	mockDB := new(MockUsersDatabase)
	service := newServiceWithMock(mockDB)

	cmd := updateuser.UpdateUserCommand{
		UserId: uuid.New(),
		Email:  "duplicate@example.com",
		Name:   "Name",
		Bio:    utils.StringPointer("Bio"),
	}

	mockDB.On("UpdateById", mock.Anything, (*sqlx.Tx)(nil), mock.AnythingOfType("*domain.User")).
		Return(true, nil)

	err := service.UpdateUser(context.Background(), cmd)
	assert.ErrorIs(t, err, errs.BadRequestError("The email requested is not available"))
}

func TestUpdateUser_UserNotFound(t *testing.T) {
	mockDB := new(MockUsersDatabase)
	service := newServiceWithMock(mockDB)

	cmd := updateuser.UpdateUserCommand{
		UserId: uuid.New(),
		Email:  "missing@example.com",
		Name:   "Missing",
		Bio:    utils.StringPointer("Bio"),
	}

	mockDB.On("UpdateById", mock.Anything, (*sqlx.Tx)(nil), mock.AnythingOfType("*domain.User")).
		Return(false, errors.New("user not found"))

	err := service.UpdateUser(context.Background(), cmd)
	assert.ErrorIs(t, err, errs.NotFoundError("User not found"))
}

func TestUpdateUser_DatabaseError(t *testing.T) {
	mockDB := new(MockUsersDatabase)
	service := newServiceWithMock(mockDB)

	cmd := updateuser.UpdateUserCommand{
		UserId: uuid.New(),
		Email:  "fail@example.com",
		Name:   "Failure",
		Bio:    utils.StringPointer("Bio"),
	}

	mockDB.On("UpdateById", mock.Anything, (*sqlx.Tx)(nil), mock.AnythingOfType("*domain.User")).
		Return(false, errors.New("db timeout"))

	err := service.UpdateUser(context.Background(), cmd)
	assert.EqualError(t, err, "db timeout")
}
