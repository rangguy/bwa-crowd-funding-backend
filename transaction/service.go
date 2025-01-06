package transaction

type Service interface {
	GetTransactionsByCampaignID(campaignID int) ([]Transaction, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetTransactionsByCampaignID(campaignID int) ([]Transaction, error) {
	transactions, err := s.repo.GetByCampaignID(campaignID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
