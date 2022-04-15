package service

import (
	"time"

	"github.com/artomsopun/mortgage/mortgage-api/internal/repository"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/auth"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/hash"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
)

type UserInfo struct {
	ID    types.BinaryUUID
	Nick  string
	Email string
}

type UserInputSigUp struct {
	Nick      string
	Email     string
	Passwords Passwords
}

type UserInputSigIn struct {
	Login    string
	Password string
}

type PasswordConfirm struct {
	UserID      types.BinaryUUID
	OldPassword string
	Passwords   Passwords
}

type Passwords struct {
	Password string
	Confirm  string
}

type Tokens struct {
	AccessToken  Token
	RefreshToken Token
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type Auths interface {
	SignUp(input UserInputSigUp) error
	SignIn(input UserInputSigIn) (Tokens, error)
	RefreshTokens(refresh Token) (Tokens, error)
}

type Profiles interface {
	GetProfile(userID types.BinaryUUID) (UserInfo, error)
	ChangePassword(confirm PasswordConfirm) error
	DeleteProfile(userID types.BinaryUUID) error
}

type Services struct {
	Auths    Auths
	Profiles Profiles
}

type Deps struct {
	Repos           *repository.Repositories
	Hasher          hash.PasswordHasher
	AuthManager     auth.AuthManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	authsService := NewAuthsService(deps.Repos.Users, deps.Repos.Sessions, deps.Hasher, deps.AuthManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	profilesService := NewProfilesService(deps.Repos.Users, deps.Hasher)

	return &Services{
		Auths:    authsService,
		Profiles: profilesService,
	}
}
