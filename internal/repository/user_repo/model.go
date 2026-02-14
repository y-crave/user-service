package user_repo

import (
	"time"

	"github.com/google/uuid"
)

type GormUser struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RegistrationState string     `gorm:"type:reg_state_enum..." `
	Name              string     `gorm:"type:varchar(100);not null"`
	BIO               string     `gorm:"type:varchar(100);not null"`
	Birthdate         time.Time  `gorm:"type:date;not null"`
	Sex               string     `gorm:"type:sex;not null"`
	Location          string     `gorm:"type:varchar(100);not null"`
	Height            int        `gorm:"type:int"`
	CreatedAt         time.Time  `gorm:"type:date;not null"`
	UpdatedAt         time.Time  `gorm:"type:date;not null"`
	DeletedAt         *time.Time // мне кажется это нам не надо
}

func (GormUser) TableName() string {
	return "users"
}

type GormPicture struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	IsMain   bool      `gorm:"not null"`
	Approved bool      `gorm:"not null"`
	Path     string    `gorm:"type:varchar(255);not null"`
}

func (GormPicture) TableName() string {
	return "pictures"
}

type GormTarget struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"type:varchar(100);not null"`
}

func (GormTarget) TableName() string {
	return "targets"
}

type GormTag struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Volume   int       `gorm:"type:int;not null"`
	Category string    `gorm:"type:varchar(100);not null"`
}

func (GormTag) TableName() string {
	return "tags"
}

type GormCategory struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string    `gorm:"type:varchar(100);not null"`
	Volume int       `gorm:"type:int;not null"`
}

func (GormCategory) TableName() string {
	return "category"
}
