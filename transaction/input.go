package transaction

import "bwastartup/user"

type GetTransactionByCampaignIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
