package service

import (
	"errors"
	"time"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"github.com/artomsopun/mortgage/mortgage-api/internal/repository"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/auth"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/hash"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
)

type AuthService struct {
	repoUser    repository.Users
	repoSession repository.Sessions
	hasher      hash.PasswordHasher
	AuthManager auth.AuthManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthsService(repoUser repository.Users, repoSession repository.Sessions,
	hasher hash.PasswordHasher, authManager auth.AuthManager,
	accessTTL, refreshTTL time.Duration) *AuthService {
	return &AuthService{
		repoUser:        repoUser,
		repoSession:     repoSession,
		hasher:          hasher,
		AuthManager:     authManager,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

func (s *AuthService) SignUp(input UserInputSigUp) error {
	if input.Passwords.Password != input.Passwords.Confirm {
		return errors.New("passwords mismatch")
	}
	passwordHash, err := s.hasher.Hash(input.Passwords.Confirm)
	if err != nil {
		return err
	}

	user := domain.User{
		Nick:     input.Nick,
		Email:    input.Email,
		Password: passwordHash,
	}

	if err := s.repoUser.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) SignIn(input UserInputSigIn) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := s.repoUser.GetByCredentials(input.Login, passwordHash)
	if err != nil {
		return Tokens{}, errors.New("user not found")
	}

	return s.createSession(user.ID)
}

func (s *AuthService) RefreshTokens(refreshToken Token) (Tokens, error) {
	session, err := s.repoSession.GetRefreshToken(refreshToken.Value)
	if err != nil {
		return Tokens{}, errors.New("wrong refresh token")
	}

	//check refresh expiring
	if session.ExpiresAt.Before(refreshToken.ExpiresAt) {
		return Tokens{}, errors.New("wrong refresh token")
	}

	return s.createSession(session.UserID)
}

func (s *AuthService) createSession(userID types.BinaryUUID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken.Value, err = s.AuthManager.NewJWT(userID.String(), s.accessTokenTTL)
	if err != nil {
		return Tokens{}, err
	}
	res.AccessToken.ExpiresAt = time.Now().Add(s.accessTokenTTL)

	res.RefreshToken.Value, err = s.AuthManager.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}
	res.RefreshToken.ExpiresAt = time.Now().Add(s.refreshTokenTTL)

	err = s.repoSession.SetSession(domain.Session{
		RefreshToken: res.RefreshToken.Value,
		ExpiresAt:    res.RefreshToken.ExpiresAt,
		UserID:       userID,
	})
	if err != nil {
		return Tokens{}, err
	}

	return res, nil
}
