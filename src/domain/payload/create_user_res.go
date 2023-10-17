package payload

import (
	"time"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
)

type CreateUserResponse struct {
	GUID        string    `json:"guid"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

func CreateUserResponses(user model.User) (res CreateUserResponse) {
	res = CreateUserResponse{
		GUID:        user.GUID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		CreatedBy:   user.CreatedBy,
	}

	return
}
