package domain

import (
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bank struct {
	ID types.BinaryUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`

	Title      string  `gorm:"column:title"`
	Rate       float32 `gorm:"column:rate"`
	MaxLoan    uint    `gorm:"column:max_loan"`
	MinPayment uint    `gorm:"column:min_payment"`
	LoanTerm   uint    `gorm:"column:loan_term"`

	UserID types.BinaryUUID `gorm:"column:user_id;"`
}

func (b *Bank) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	b.ID = types.BinaryUUID(id)
	return err
}
