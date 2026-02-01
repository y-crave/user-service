package user_repo

import (
	"user-service/internal/domain"
)

func toDomainUser(user GormUser) domain.User {
	return domain.User{
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
