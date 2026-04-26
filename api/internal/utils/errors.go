package utils

import "errors"

// Define custom error variables for common error scenarios in the application.
// These errors can be used throughout the codebase to provide consistent error handling and messaging.
var (
	ErrPortRequired       = errors.New("PORT environment variable is required")
	ErrInvalidPort        = errors.New("invalid PORT value")
	ErrDBHostRequired     = errors.New("DBHOST environment variable is required")
	ErrDBUserRequired     = errors.New("DBUSER environment variable is required")
	ErrDBPasswordRequired = errors.New("DBPASSWORD environment variable is required")
	ErrDBNameRequired     = errors.New("DBNAME environment variable is required")
	ErrDBPortRequired     = errors.New("DBPORT environment variable is required")
	ErrInvalidDBPort      = errors.New("invalid DBPORT value")
	ErrSellerIDNotFound   = errors.New("seller not found")
	ErrInvalidID          = errors.New("invalid ID parameter")
)
