package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int               `gorm:"primaryKey;autoIncrement"`
	UserID     int               `gorm:"column:user_id"`
	CampaignID int               `gorm:"column:campaign_id"`
	Amount     int               `gorm:"column:amount;type:integer"`
	Status     string            `gorm:"column:status;type:varchar"`
	Code       string            `gorm:"column:code;type:varchar"`
	PaymentURL string            `gorm:"column:payment_url;type:varchar"`
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
	UpdatedAt  time.Time         `gorm:"autoUpdateTime"`
	User       user.User         `gorm:"foreignKey:UserID"`
	Campaign   campaign.Campaign `gorm:"foreignKey:CampaignID"`
}
