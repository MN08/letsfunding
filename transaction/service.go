package transaction

import (
	"errors"
	"letsfunding/campaign"
)

type (
	service struct {
		repository         Repository
		campaignRepository campaign.Repository
	}

	Service interface {
		GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
		GetTransactionByUserID(userID int) ([]Transaction, error)
	}
)

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("only owner can access")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
