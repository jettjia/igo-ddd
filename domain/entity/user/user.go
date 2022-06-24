package entity

import "time"

type User struct {
	ID        uint64     `json:"id"`
	Passport  string     `json:"passport"`
	Password  string     `json:"Password"`
	Nickname  string     `json:"nickname"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
