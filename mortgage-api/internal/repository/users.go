package repository

import (
	"errors"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetById(userID types.BinaryUUID) (domain.User, error) {
	user := domain.User{}
	if err := r.db.First(&user, userID).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UsersRepo) GetByCredentials(login, password string) (domain.User, error) {
	user := domain.User{}
	if err := r.db.Where("(nick = ? OR email = ?) AND password = ?", login, login, password).First(&user).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UsersRepo) Create(user domain.User) error {
	err := r.db.Where("email = ?", user.Email).First(&domain.User{}).Error
	if err == nil {
		return errors.New("user already exists")
	}
	err = r.db.Where("nick = ?", user.Nick).First(&domain.User{}).Error
	if err == nil {
		return errors.New("user already exists")
	}
	r.db.Create(&user)
	return nil
}

func (r *UsersRepo) ChangePassword(userID types.BinaryUUID, password string) error {
	err := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("password", password).Error
	return err
}

func (r *UsersRepo) Delete(userID types.BinaryUUID) error {
	err := r.db.Delete(&domain.User{}, userID).Error
	return err
}
