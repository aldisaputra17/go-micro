package dto

type CreateUserParam struct {
	GUID        string `db:"guid"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	CreatedBy   string `db:"created_by"`
}
