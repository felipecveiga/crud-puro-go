package errs

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidPayload       = errors.New("invalid payload")
	ErrUserID               = errors.New("user ID error")
	ErrInvalidMethodRequest = errors.New("invalid HTTP method for request")
	ErrBodyRequest          = errors.New("error creating account, mandatory filling of name, email and phone number")
	ErrInsertUserDatabase   = errors.New("error when inserting into the database")
	ErrConvertIDObjectID    = errors.New("error converting ID")
	ErrSearchUser           = errors.New("error searching for user")
	
)
