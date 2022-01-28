package models

import "time"

type User struct {
	ID        uint
	FullName  string
	Email     string
	CreacteAt time.Time
	UpdateAt time.Time
}

