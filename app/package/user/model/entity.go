package model

import (
	"time"
)

type User struct {
	ID           int64      `gorm:"primaryKey;autoIncrement;not null"`
	Name         string     `gorm:"size:255;not null"`
	Email        string     `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string     `gorm:"size:255;not null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
}

func (u *User) ToRegisterResponse() RegisterResponse {
	return RegisterResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
