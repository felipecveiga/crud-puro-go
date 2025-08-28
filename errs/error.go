package errs

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserID           = errors.New("user ID error")
	ErrUserInsertFailed = errors.New("failed to insert user into database")
	ErrUserSearchFailed = errors.New("failed to search for user")

	ErrInvalidPayload        = errors.New("invalid payload")
	ErrInvalidHTTPMethod     = errors.New("invalid HTTP method")
	ErrMissingRequiredFields = errors.New("missing required fields: name, email, and phone number")
	ErrInvalidObjectID       = errors.New("invalid object ID format")
)
