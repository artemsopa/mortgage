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

type Bank struct {
	ID types.BinaryUUID

	Title      string
	Rate       float64
	MaxLoan    uint
	MinPayment uint
	LoanTerm   uint

	UserID types.BinaryUUID
}

type CalculateInput struct {
	Loan    uint
	Payment uint
	BankID  types.BinaryUUID
}

type Calculator struct {
	Principal float64
	Rate      float64
	Months    int
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

type Banks interface {
	GetAllBanks() ([]Bank, error)
	GetBanksByUserID(userID types.BinaryUUID) ([]Bank, error)
	CreateBank(bank Bank) error
	UpdateBank(bank Bank) error
	DeleteBank(userID, bankID types.BinaryUUID) error
	CalculateMortgage(input CalculateInput) string
}

type Services struct {
	Auths    Auths
	Profiles Profiles
	Banks    Banks
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
	banksService := NewBanksService(deps.Repos.Banks)

	return &Services{
		Auths:    authsService,
		Profiles: profilesService,
		Banks:    banksService,
	}
}
