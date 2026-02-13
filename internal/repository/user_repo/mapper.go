package user_repo

import (
	"user-service/internal/domain"

	"github.com/google/uuid"
)

func toDomainUser(u *GormUser) *domain.User {
	return &domain.User{
		ID:                uuid.MustParse(u.ID),
		Email:             user.Email,
		RegistrationState: user.RegistrationState,
		Name:              user.Name,
		BIO:               user.BIO,
		Birthdate:         user.Birthdate,
		Sex:               user.Sex,
		Location:          user.Location,
		Height:            user.Height,
	}
}

func toGormUser(user domain.User) GormUser {
	return GormUser{
		ID:                user.ID,
		Email:             user.Email,
		RegistrationState: user.RegistrationState,
		Name:              user.Name,
		BIO:               user.BIO,
		Birthdate:         user.Birthdate,
		Sex:               user.Sex,
		Location:          user.Location,
		Height:            user.Height,
	}
}
