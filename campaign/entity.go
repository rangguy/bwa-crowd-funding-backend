package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               int             `gorm:"primaryKey;autoIncrement"`
	UserID           int             `gorm:"column:user_id"`
	Name             string          `gorm:"type:varchar"`
	ShortDescription string          `gorm:"column:short_description;type:varchar"`
	Description      string          `gorm:"type:text"`
	Perks            string          `gorm:"type:text"`
	BackerCount      int             `gorm:"column:backer_count;type:integer"`
	GoalAmount       int             `gorm:"column:goal_amount;type:integer"`
	CurrentAmount    int             `gorm:"column:current_amount;type:integer"`
	Slug             string          `gorm:"type:varchar"`
	CreatedAt        time.Time       `gorm:"autoCreateTime"`
	UpdatedAt        time.Time       `gorm:"autoUpdateTime"`
	CampaignImages   []CampaignImage `gorm:"foreignKey:CampaignID"`
	User             user.User       `gorm:"foreignKey:UserID"`
}

type CampaignImage struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	CampaignID int       `gorm:"column:campaign_id"`
	FileName   string    `gorm:"column:file_name;type:varchar"`
	IsPrimary  bool      `gorm:"column:is_primary;type:boolean"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
