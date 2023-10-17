package payload

import (
	"time"

	"github.com/aldisaputra17/go-micro/src/domain/repository/model"
)

type UserResponse struct {
	ID          int64     `json:"id"`
	GUID        string    `db:"guid"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   *string   `json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   *string   `json:"deleted_by"`
}

func ToReadUserResponse(user model.User) (res UserResponse) {
	res = UserResponse{
		ID:          user.ID,
		GUID:        user.GUID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		CreatedBy:   user.CreatedBy,
	}

	if user.UpdatedAt.Valid {
		res.UpdatedAt = user.UpdatedAt.Time
	}

	if user.UpdatedBy.Valid {
		res.UpdatedBy = &user.UpdatedBy.String
	}

	if user.DeletedAt.Valid {
		res.DeletedAt = user.DeletedAt.Time
	}

	if user.DeletedBy.Valid {
		res.DeletedBy = &user.DeletedBy.String
	}

	return
}

type LoginResponse struct {
	GUID                  string       `json:"guid"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiredAt  time.Time    `json:"access_token_expired_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expired_at"`
	CreatedAt             time.Time    `json:"created_at"`
	User                  UserResponse `json:"user"`
}
