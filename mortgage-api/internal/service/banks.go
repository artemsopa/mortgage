package service

import (
	"errors"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"github.com/artomsopun/mortgage/mortgage-api/internal/repository"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
)

type BanksService struct {
	repoBanks repository.Banks
}

func NewBanksService(repoBanks repository.Banks) *BanksService {
	return &BanksService{
		repoBanks: repoBanks,
	}
}

func (s *BanksService) GetAllBanks() ([]Bank, error) {
	banksRepo, err := s.repoBanks.GetAllBanks()
	if err != nil {
		return []Bank{}, err
	}
	var banks []Bank
	for _, bank := range banksRepo {
		banks = append(banks, Bank{
			ID:         bank.ID,
			Title:      bank.Title,
			Rate:       bank.Rate,
			MaxLoan:    bank.MaxLoan,
			MinPayment: bank.MinPayment,
			LoanTerm:   bank.LoanTerm,
			UserID:     bank.UserID,
		})
	}
	return banks, nil
}

func (s *BanksService) GetBanksByUserID(userID types.BinaryUUID) ([]Bank, error) {
	banksRepo, err := s.repoBanks.GetBanksByUserID(userID)
	if err != nil {
		return []Bank{}, err
	}
	var banks []Bank
	for _, bank := range banksRepo {
		banks = append(banks, Bank{
			ID:         bank.ID,
			Title:      bank.Title,
			Rate:       bank.Rate,
			MaxLoan:    bank.MaxLoan,
			MinPayment: bank.MinPayment,
			LoanTerm:   bank.LoanTerm,
			UserID:     bank.UserID,
		})
	}
	return banks, nil
}

func (s *BanksService) CreateBank(bank Bank) error {
	err := s.repoBanks.Create(domain.Bank{
		Title:      bank.Title,
		Rate:       bank.Rate,
		MaxLoan:    bank.MaxLoan,
		MinPayment: bank.MinPayment,
		LoanTerm:   bank.LoanTerm,
		UserID:     bank.UserID,
	})
	return err
}

func (s *BanksService) UpdateBank(bank Bank) error {
	bankRepo, err := s.repoBanks.GetById(bank.ID)
	if err != nil {
		return err
	}
	if bank.UserID != bankRepo.UserID {
		return errors.New("you can update your banks only")
	}
	err = s.repoBanks.Update(domain.Bank{
		ID:         bank.ID,
		Title:      bank.Title,
		Rate:       bank.Rate,
		MaxLoan:    bank.MaxLoan,
		MinPayment: bank.MinPayment,
		LoanTerm:   bank.LoanTerm,
		UserID:     bank.UserID,
	})
	return err
}

func (s *BanksService) DeleteBank(userID, bankID types.BinaryUUID) error {
	err := s.repoBanks.Delete(userID, bankID)
	return err
}

//func (s *BanksService) CalculateMortgage(input CalculateInput) (string, error) {}
