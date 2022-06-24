package po

type UserLog struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserId uint64 `gorm:"bigint" json:"user_id"`
	Log    string `gorm:"size:64;not null;" json:"log"`
}
