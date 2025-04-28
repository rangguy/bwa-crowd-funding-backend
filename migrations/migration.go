package migrations

import (
	"bwastartup/campaign"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{}, &transaction.Transaction{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
