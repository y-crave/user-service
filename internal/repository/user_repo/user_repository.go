package user_repo

import (
	"context"
	"errors"
	"fmt"
	"user-service/internal/domain"
	"user-service/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) service.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user domain.User) error {
	gormUser := toGormUser(user)
	if err := r.db.WithContext(ctx).Create(&gormUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrUserAlreadyExists
		}

		return fmt.Errorf("%w, gorm error: %v", domain.ErrInternalError, err)
	}
	return nil
}

func (r *userRepo) Read(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	var gormUser GormUser

	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&gormUser).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{},fmt.Errorf("%w, gorm error: %v", domain.ErrInternalError, err)
	return toDomainUser(gormUser), nil
}
