package ddd

import "time"

type Model struct {
	CreatedAt time.Time `db:"created_at"`
}

type AuditableModel struct {
	Model
	UpdatedAt *time.Time `db:"updated_at"`
}
