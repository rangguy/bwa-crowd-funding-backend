package transaction

import "bwastartup/user"

type GetTransactionByCampaignIDInput struct {
	ID   int `uri:"id"`
	User user.User
}
