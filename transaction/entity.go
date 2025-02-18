package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionMigrate struct {
	ID         int               `gorm:"primary_key;AUTO_INCREMENT"`
	UserID     int               `gorm:"column:user_id"`
	CampaignID int               `gorm:"column:campaign_id"`
	Amount     int               `gorm:"column:amount;type:integer"`
	Status     string            `gorm:"column:status;type:varchar"`
	Code       string            `gorm:"column:code;type:varchar"`
	CreatedAt  time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	User       user.User         `gorm:"foreignKey:UserID"`
	Campaign   campaign.Campaign `gorm:"foreignKey:CampaignID"`
}
