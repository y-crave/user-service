package service

import (
	"context"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

type UserService interface {
}

type UserRepo interface {
	Create(ctx context.Context, user domain.User) error
	Read(ctx context.Context, userID uuid.UUID) error
}
