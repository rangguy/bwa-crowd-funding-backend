package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(campaignID GetTransactionByCampaignIDInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

type service struct {
	repo         Repository
	campaignRepo campaign.Repository
}

func NewService(repo Repository, campaignRepo campaign.Repository) Service {
	return &service{repo, campaignRepo}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionByCampaignIDInput) ([]Transaction, error) {
	campaign, err := s.campaignRepo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("user is not the owner of the campaign")
	}

	transactions, err := s.repo.GetByCampaignID(input.ID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
