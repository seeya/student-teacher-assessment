package models

import (
	"time"
)

type Teacher struct {
	ID        int64     `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
}
