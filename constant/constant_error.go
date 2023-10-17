package constant

import "errors"

// error message.
const (
	MsgHeaderTokenNotFound        = "header `authorization` not found"
	MsgHeaderRefreshTokenNotFound = "header `refresh-token` not found"
	MsgHeaderTokenInvalid         = "invalid header token"
	MsgHeaderRefreshTokenInvalid  = "invalid header refresh token"
	MsgTokenUnauthorized          = "unauthorized token"
	MsgRefreshTokenUnauthorized   = "unauthorized refresh token"
	MsgTokenInvalid               = "invalid token"
	MsgRefreshTokenInvalid        = "invalid refresh token"
	MsgTokenExpired               = "expired token"
	MsgRefreshTokenExpired        = "expired refresh token"

	MsgForbiddenLevel      = "your level is not allowed to access this resource"
	MsgForbiddenPermission = "your permission is not allowed to access this resource"

	MsgForgotPasswordLinkInvalid = "invalid forgot password link"
	MsgForgotPasswordLinkExpired = "expired forgot password link"

	MsgErrStorageConn = "failed to connect with storage"
)

// error interface.
var (
	// 400.
	ErrFailedParseRequest   = errors.New("failed to parse request")
	ErrValidationFailed     = errors.New("validation failed")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrpasswordDoesNotMatch = errors.New("password does not match")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrDataAlreadyExists    = errors.New("data already exists")
	ErrFailedLogin          = errors.New("email and password do not match")
	ErrDataFailedCreated    = errors.New("failed create data")
	ErrDataUpdated          = errors.New("failed update data")
	ErrPhoneNumberInvalid   = errors.New("phone number not valid")
	ErrFileNotFoundRequest  = errors.New("failed to parse request file")
	ErrFailedDelete         = errors.New("failed delete data")
	ErrFileNotSupport       = errors.New("file not supported")

	// 401.
	ErrHeaderTokenNotFound = errors.New("header authorization not found")
	ErrHeaderTokenInvalid  = errors.New("invalid header token")
	ErrTokenUnauthorized   = errors.New("unauthorized token")
	ErrTokenInvalid        = errors.New("invalid token")
	ErrTokenExpired        = errors.New("expired token")

	// 403.
	ErrForbiddenRole       = errors.New("your role is not allowed to access this resource")
	ErrForbiddenPermission = errors.New("your permission is not allowed to access this resource")

	// 404.
	ErrAccountNotFound = errors.New("account not found")
	ErrDataNotFound    = errors.New("data not found")

	// 500.
	ErrUnknownSource          = errors.New("an error occurred, please try again later")
	ErrAccountNotHavePassword = errors.New("account does not have password")
	ErrTaskFailed             = errors.New("failed to do the task")
	ErrFileUpload             = errors.New("failed to upload file")
	ErrFileOpenFailed         = errors.New("failed open file")
)