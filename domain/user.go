package domain

import "time"

type User struct {
	UserId    string
	Name      string
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User
