package model

const (
	ErrEmailAlreadyExists = "email already exists"
	ErrPasswordHashing    = "failed to hash password"
	ErrDatabase           = "failed to connect to database"
	ErrUserCreation       = "failed to create user"
	ErrInvalidRequest     = "invalid request"
	ErrUnauthorizedAccess = "unauthorized access"
	ErrNotFound           = "not found"
	ErrInternalError      = "internal server error"
	ErrInvalidPassword    = "invalid password"
	ErrMissingFields      = "missing required fields"
	ErrEmailFormat        = "invalid email format"
	ErrPasswordLength     = "password must be at least 8 characters long"
	ErrPasswordComplexity = "password must contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	ErrTokenGeneration    = "failed to generate token"
)

type RegisterResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
