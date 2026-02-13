package user_repo

import (
	"github.com/google/uuid"
	"time"
	"user-service/internal/domain"
)

type GormUser struct {
	ID                uuid.UUID
	Email             string
	RegistrationState string `gorm:"type:reg_state_enum..." `
	Name              string
	BIO               string
	Birthdate         time.Time
	Sex               string
	Location          string
	Height            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func (GormUser) TableName() string {
	return "user"
}

type GormPicture struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	IsMain   bool
	Approved bool
	Path     string
}

func (GormPicture) TableName() string {
	return "picture"
}

type GormTarget struct {
	ID   uuid.UUID
	Name string
}

func (GormTarget) TableName() string {
	return "target"
}

type GormTag struct {
	ID       uuid.UUID
	Name     string
	Volume   int
	Category string
}

func (GormTag) TableName() string {
	return "tag"
}

type GormCategory struct {
	ID     uuid.UUID
	Name   string
	Volume int
}

func (GormCategory) TableName() string {
	return "category"
}
