package params

import (
	"final-project/entity"
	"final-project/helper"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,numeric,min=8"`
}
type ResponseUser struct {
	ID       uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (c *CreateUserRequest) ParseToModel() *entity.User {
	passwordHash, _ := helper.HashPassword(c.Password)

	return &entity.User{
		Username: c.Username,
		Email:    c.Email,
		Password: passwordHash,
		Age:      c.Age,
	}
}

func ResponseFromEntity(user *entity.User) *ResponseUser {
	return &ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
