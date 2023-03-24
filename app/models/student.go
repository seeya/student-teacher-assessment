package models

import (
	"time"
)

type Student struct {
	ID          int64     `db:"id" json:"id" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Email       string    `db:"email" json:"email" validate:"required,email"`
	IsSuspended bool      `db:"is_suspended" json:"is_suspended"`
}
