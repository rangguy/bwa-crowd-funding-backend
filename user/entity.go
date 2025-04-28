package user

import "time"

type User struct {
	ID             int       `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"type:varchar"`
	Occupation     string    `gorm:"type:varchar"`
	Email          string    `gorm:"type:varchar"`
	PasswordHash   string    `gorm:"column:password_hash;type:varchar"`
	AvatarFileName string    `gorm:"column:avatar_file_name;type:varchar"`
	Role           string    `gorm:"type:varchar"`
	Token          string    `gorm:"type:varchar"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
