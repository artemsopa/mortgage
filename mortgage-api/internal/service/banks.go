package service

import (
	"errors"
	"math"
	"strconv"

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
	bankRepo, err := s.repoBanks.GetById(bankID)
	if err != nil {
		return err
	}
	if userID != bankRepo.UserID {
		return errors.New("you can delete your banks only")
	}
	err = s.repoBanks.Delete(userID, bankID)
	return err
}

func (s *BanksService) CalculateMortgage(input CalculateInput) string {
	bankRepo, err := s.repoBanks.GetById(input.BankID)
	if err != nil {
		return "Bank not found"
	}
	if input.Loan > bankRepo.MaxLoan && input.Payment < bankRepo.MinPayment {
		return "False initial loan & down payment"
	}
	if input.Loan > bankRepo.MaxLoan {
		return "Initial loan higher then max"
	}
	if input.Payment < bankRepo.MinPayment {
		return "Down payment lower then min"
	}
	c := Calculator{
		Loan:   float64(input.Loan),
		Down:   float64(input.Payment),
		Rate:   bankRepo.Rate,
		Months: int(bankRepo.LoanTerm),
	}
	result := (c.Loan*(c.Rate/12)*(math.Pow((1+(c.Rate/12)), 5)))/math.Pow((1+(c.Rate/12)), 5) - 1
	return strconv.Itoa(int(result))
}
