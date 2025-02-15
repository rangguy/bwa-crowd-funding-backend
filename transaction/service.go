package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(campaignID GetTransactionByCampaignIDInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repo           Repository
	campaignRepo   campaign.Repository
	paymentService payment.Service
}

func NewService(repo Repository, campaignRepo campaign.Repository, paymentService payment.Service) Service {
	return &service{repo, campaignRepo, paymentService}
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

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		Amount:     input.Amount,
		CampaignID: input.CampaignID,
		UserID:     input.User.ID,
		Status:     "pending",
	}

	newTransaction, err := s.repo.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repo.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
