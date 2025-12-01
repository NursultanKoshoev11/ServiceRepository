package models

import "time"

type User struct {
	ID        int64     `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	RoleID    RoleType  `db:"role_id" json:"role_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Profile struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Avatar    string    `db:"avatar" json:"avatar,omitempty"`
	Bio       string    `db:"bio" json:"bio,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Role struct {
	ID   int16  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type RoleType int16

const (
	RoleUser      RoleType = iota // 0
	RoleModerator                 // 1
	RoleAdmin                     // 2
)
