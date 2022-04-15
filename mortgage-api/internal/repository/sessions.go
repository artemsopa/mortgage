package repository

import (
	"errors"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"gorm.io/gorm"
)

type SessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *SessionsRepo {
	return &SessionsRepo{
		db: db,
	}
}

func (r *SessionsRepo) GetRefreshToken(refreshToken string) (domain.Session, error) {
	session := domain.Session{}
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&session).Error; err != nil {
		return domain.Session{}, errors.New("session not found")
	}
	return session, nil
}

func (r *SessionsRepo) SetSession(session domain.Session) error {
	err := r.db.Where("user_id = ?", session.UserID).First(&domain.Session{}).Error
	if err != nil {
		r.db.Create(&session)
	} else {
		r.db.Model(domain.Session{}).Where("user_id = ?", session.UserID).Updates(&session)
	}
	return nil
}
