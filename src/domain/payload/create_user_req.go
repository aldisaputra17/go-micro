package payload

import (
	"context"

	"github.com/aldisaputra17/go-micro/helper"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	GUID        string `json:"guid"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedBy   string `json:"created_by"`
}

func (req *CreateUserRequest) ToParam() (user dto.CreateUserParam) {
	hashPassword, err := helper.GeneratePassword(req.Password)
	if err != nil {
		log.FromCtx(context.Background()).Error(err, "failed hash password")
		return
	}

	user = dto.CreateUserParam{
		GUID:        uuid.NewString(),
		Email:       req.Email,
		Password:    hashPassword,
		PhoneNumber: req.Password,
		CreatedBy:   req.CreatedBy,
	}

	return
}
