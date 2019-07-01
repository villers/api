package datamodel

import "time"

type User struct {
	ID        int64      `json:"id"`
	Login     string     `json:"login"`
	Password  string     `json:"password"`
	Email     string     `json:"email" validate:"required,email"`
	Companies []*Company `json:"compagnies" validate:"required,dive,required"`
	CreatedAt time.Time  `json:"created_at"`
}
