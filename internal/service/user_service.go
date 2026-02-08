package service

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return &userService{userRepo: userRepo}
}
