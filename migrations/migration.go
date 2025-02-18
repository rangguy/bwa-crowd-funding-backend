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
	err = db.AutoMigrate(&user.UserMigrate{}, &campaign.CampaignMigrate{}, &campaign.CampaignImageMigrate{}, &transaction.TransactionMigrate{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
