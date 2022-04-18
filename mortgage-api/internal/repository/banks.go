package repository

import (
	"errors"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"gorm.io/gorm"
)

type BanksRepo struct {
	db *gorm.DB
}

func NewBanksRepo(db *gorm.DB) *BanksRepo {
	return &BanksRepo{
		db: db,
	}
}

func (r *BanksRepo) GetById(bankID types.BinaryUUID) (domain.Bank, error) {
	bank := domain.Bank{}
	if err := r.db.Where("id = ?", bankID).First(&bank).Error; err != nil {
		return domain.Bank{}, errors.New("bank not found")
	}
	return bank, nil
}

func (r *BanksRepo) GetAllBanks() ([]domain.Bank, error) {
	var banks []domain.Bank
	if err := r.db.Find(&banks).Error; err != nil {
		return []domain.Bank{}, err
	}
	return banks, nil
}

func (r *BanksRepo) GetBanksByUserID(userID types.BinaryUUID) ([]domain.Bank, error) {
	var banks []domain.Bank
	if err := r.db.Where("user_id = ?", userID).Find(&banks).Error; err != nil {
		return []domain.Bank{}, err
	}
	return banks, nil
}

func (r *BanksRepo) Create(bank domain.Bank) error {
	err := r.db.Where("title = ?", bank.Title).First(&domain.Bank{}).Error
	if err == nil {
		return errors.New("bank title already exist")
	}
	err = r.db.Create(&bank).Error
	return err
}

func (r *BanksRepo) Update(bank domain.Bank) error {
	err := r.db.Where("id != ? AND title = ?", bank.ID, bank.Title).First(&domain.Bank{}).Error
	if err == nil {
		return errors.New("bank title already exist")
	}
	err = r.db.Save(&bank).Error
	return err
}

func (r *BanksRepo) Delete(userID, bankID types.BinaryUUID) error {
	err := r.db.Where("id = ? AND user_id = ?", bankID, userID).Delete(&domain.Bank{}).Error
	return err
}
