package transaction

import "time"

type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	formatter := CampaignTransactionsFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}

	var transactionsFormatter []CampaignTransactionsFormatter

	for _, transaction := range transactions {
		transactionsFormatter = append(transactionsFormatter, FormatCampaignTransaction(transaction))
	}

	return transactionsFormatter
}
