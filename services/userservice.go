package services

import (
	"final-project/entity"
	"final-project/params"
	"final-project/repository"
	"fmt"
)

type UserServices struct {
	userRepo repository.UserRepository
}

func NewCarServices(userRepo repository.UserRepository) *UserServices {
	return &UserServices{
		userRepo: userRepo,
	}
}

func (u *UserServices) RegisterUser(req *params.CreateUserRequest) *entity.Response {
	user := req.ParseToModel()
	err := u.userRepo.RegisterUser(user)
	if err != nil {
		msg := fmt.Sprintf("there is something wrong")
		return entity.InternalServerError("INTERNAL SERVER ERROR", &msg, nil)
	}

	return entity.CreatedSuccess(nil, user)
}

func (u *UserServices) GetAllUser() *entity.Response {
	err := u.userRepo.GetAllUser()
	// fmt.Println(err)
	if err != nil {
		msg := fmt.Sprintf("User not found")
		return entity.NotFound(nil, &msg, nil)
	}

	return entity.FindSuccess(nil, err)
}
