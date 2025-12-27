package model

import "time"

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	IsActive  bool
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
