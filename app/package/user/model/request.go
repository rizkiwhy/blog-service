package model

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (r *RegisterRequest) HashPassword() (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("[HashPassword] Failed to hash password")
		return
	}

	r.Password = string(hashedPassword)

	return
}

func (r *RegisterRequest) ToUser() User {
	return User{
		Name:         r.Name,
		Email:        r.Email,
		PasswordHash: r.Password,
	}
}
