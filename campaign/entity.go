package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CampaignMigrate struct {
	ID               int       `gorm:"primary_key;AUTO_INCREMENT"`
	UserID           int       `gorm:"column:user_id"`
	Name             string    `gorm:"column:name;type:varchar"`
	ShortDescription string    `gorm:"column:short_description;type:varchar"`
	Description      string    `gorm:"column:description;type:text"`
	GoalAmount       int       `gorm:"column:goal_amount;type:integer"`
	CurrentAmount    int       `gorm:"column:current_amount;type:integer"`
	Perks            string    `gorm:"column:perks;type:text"`
	BackerCount      int       `gorm:"column:backer_count;type:integer"`
	Slug             string    `gorm:"column:slug;type:varchar"`
	User             user.User `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CampaignImageMigrate struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT"`
	CampaignID int       `gorm:"column:campaign_id"`
	FileName   string    `gorm:"column:file_name;type:varchar"`
	IsPrimary  bool      `gorm:"column:is_primary;type:boolean"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
