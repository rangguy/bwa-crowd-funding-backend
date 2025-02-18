package user

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserMigrate struct {
	ID             int       `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string    `gorm:"column:name;type:varchar"`
	Occupation     string    `gorm:"column:occupation;type:varchar"`
	Email          string    `gorm:"column:email;type:varchar"`
	PasswordHash   string    `gorm:"column:password_hash;type:varchar"`
	AvatarFileName string    `gorm:"column:avatar_file_name;type:varchar"`
	Role           string    `gorm:"column:role;type:varchar"`
	Token          string    `gorm:"column:token;type:varchar"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
