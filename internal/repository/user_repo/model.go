package user_repo

import (
	"time"

	"github.com/google/uuid"
)

type GormUser struct {
	ID                uuid.UUID
	Email             string
	RegistrationState RegStateEnum
	Name              string
	BIO               string
	Birthdate         time.Time
	Sex               SexEnum
	Location          string
	Height            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
