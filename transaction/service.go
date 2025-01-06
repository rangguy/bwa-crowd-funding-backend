package transaction

type Service interface {
	GetTransactionsByCampaignID(campaignID GetTransactionByCampaignIDInput) ([]Transaction, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionByCampaignIDInput) ([]Transaction, error) {
	transactions, err := s.repo.GetByCampaignID(input.ID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
