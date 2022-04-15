package service

import (
	"errors"

	"github.com/artomsopun/mortgage/mortgage-api/internal/repository"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/hash"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
)

type ProfilesService struct {
	repoUsers repository.Users
	hasher    hash.PasswordHasher
}

func NewProfilesService(repoUsers repository.Users, hasher hash.PasswordHasher) *ProfilesService {
	return &ProfilesService{
		repoUsers: repoUsers,
		hasher:    hasher,
	}
}

func (s *ProfilesService) GetProfile(userID types.BinaryUUID) (UserInfo, error) {
	user, err := s.repoUsers.GetById(userID)
	if err != nil {
		return UserInfo{}, err
	}
	return UserInfo{
		ID:    user.ID,
		Nick:  user.Nick,
		Email: user.Email,
	}, nil
}

func (s *ProfilesService) ChangePassword(confirm PasswordConfirm) error {
	user, err := s.repoUsers.GetById(confirm.UserID)
	if err != nil {
		return err
	}
	oldHash, err := s.hasher.Hash(confirm.OldPassword)
	if err != nil {
		return err
	}
	if user.Password != oldHash {
		return errors.New("wrong password")
	}
	if confirm.Passwords.Password != confirm.Passwords.Confirm {
		return errors.New("passwords mismatch")
	}
	passwordHash, err := s.hasher.Hash(confirm.Passwords.Confirm)
	if err != nil {
		return err
	}
	err = s.repoUsers.ChangePassword(confirm.UserID, passwordHash)
	return err
}

func (s *ProfilesService) DeleteProfile(userID types.BinaryUUID) error {
	err := s.repoUsers.Delete(userID)
	return err
}
