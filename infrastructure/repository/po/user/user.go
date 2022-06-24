package po

import "time"

type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Passport  string     `gorm:"size:45;not null;" json:"passport"`
	Password  string     `gorm:"size:45;not null" json:"Password"`
	Nickname  string     `gorm:"size:45;not null" json:"nickname"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
