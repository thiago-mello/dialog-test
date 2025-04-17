package testdb

import "github.com/google/uuid"

type TestUser struct {
	ID       uuid.UUID
	Email    string
	Password string
	Name     string
}

var User TestUser = TestUser{
	ID:       uuid.MustParse("3d0aef0a-0167-4dde-b013-75889f0ce8a3"),
	Email:    "test@example.com",
	Password: "12345678",
	Name:     "Tester",
}
