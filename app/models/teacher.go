package models

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
}
