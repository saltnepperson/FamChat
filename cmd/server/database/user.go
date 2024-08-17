package database

import "time"

type User struct {
	ID int64
	Username string
	PasswordHash string
	Email string
	CreatedBy int64
	CreatedAt time.Time
	UpdatedBy int64
	UpdatedAt time.Time
}
