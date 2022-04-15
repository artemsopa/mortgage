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

type Repositories struct {
	Users    Users
	Sessions Sessions
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db),
		Sessions: NewSessionsRepo(db),
	}
}
