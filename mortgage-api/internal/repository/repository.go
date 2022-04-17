package repository

import (
	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"gorm.io/gorm"
)

type Users interface {
	GetById(userID types.BinaryUUID) (domain.User, error)
	GetByCredentials(nickname, password string) (domain.User, error)
	Create(user domain.User) error
	ChangePassword(userID types.BinaryUUID, password string) error
	Delete(userID types.BinaryUUID) error
}

type Sessions interface {
	GetRefreshToken(refreshToken string) (domain.Session, error)
	SetSession(session domain.Session) error
}

type Banks interface {
	GetById(bankID types.BinaryUUID) (domain.Bank, error)
	GetAllBanks() ([]domain.Bank, error)
	GetBanksByUserID(userID types.BinaryUUID) ([]domain.Bank, error)
	Create(bank domain.Bank) error
	Update(bank domain.Bank) error
	Delete(userID, bankID types.BinaryUUID) error
}

type Repositories struct {
	Users    Users
	Sessions Sessions
	Banks    Banks
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Sessions: NewSessionsRepo(db),
		Banks:    NewBanksRepo(db),
	}
}
