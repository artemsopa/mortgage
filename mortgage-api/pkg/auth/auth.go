package auth

import "time"

type AuthManager interface {
	NewJWT(userID string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}
