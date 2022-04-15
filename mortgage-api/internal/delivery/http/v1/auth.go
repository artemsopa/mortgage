package v1

import (
	"net/http"
	"time"

	"github.com/artomsopun/mortgage/mortgage-api/internal/service"
	"github.com/labstack/echo/v4"
)

func (h *Handler) initAuthRoutes(api *echo.Group) {
	users := api.Group("/auth")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
		users.POST("/refresh", h.refreshSession)
		users.POST("/logout", h.logout)

	}
}

type userInputSignUp struct {
	Nick     string `json:"nick"     binding:"required,min=8,max=32"`
	Email    string `json:"email"    binding:"required,email,max=32"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	Confirm  string `json:"confirm"  binding:"required,min=8,max=32"`
}

func (h *Handler) signUp(c echo.Context) error {
	var inp userInputSignUp
	if err := c.Bind(&inp); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.Auths.SignUp(service.UserInputSigUp{
		Nick:  inp.Nick,
		Email: inp.Email,
		Passwords: service.Passwords{
			Password: inp.Password,
			Confirm:  inp.Confirm,
		},
	}); err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusCreated, "account created")
}

const (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
	DefaultPath  = "/"
)

type userInputSignIn struct {
	Login    string `json:"login"    binding:"required,min=8,max=32"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

func (h *Handler) signIn(c echo.Context) error {
	var inp userInputSignIn
	if err := c.Bind(&inp); err != nil {
		return newResponse(c, http.StatusBadRequest, "invalid input body")
	}

	res, err := h.services.Auths.SignIn(service.UserInputSigIn{
		Login:    inp.Login,
		Password: inp.Password,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	h.setCookie(c, res.AccessToken.Value, res.RefreshToken.Value, res.AccessToken.ExpiresAt, res.RefreshToken.ExpiresAt)

	return newResponse(c, http.StatusOK, "authorized, check cookies")
}

func (h *Handler) refreshSession(c echo.Context) error {
	refreshCookie, err := c.Cookie(RefreshToken)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	res, err := h.services.Auths.RefreshTokens(service.Token{
		Value:     refreshCookie.Value,
		ExpiresAt: refreshCookie.Expires,
	})
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}

	h.setCookie(c, res.AccessToken.Value, res.RefreshToken.Value, res.AccessToken.ExpiresAt, res.RefreshToken.ExpiresAt)

	return newResponse(c, http.StatusOK, "session refreshed, check cookies")
}

func (h *Handler) logout(c echo.Context) error {
	h.setCookie(c, "", "", time.Now().Add(-time.Hour), time.Now().Add(-time.Hour))

	return newResponse(c, http.StatusOK, "logout success")
}

func (h *Handler) setCookie(c echo.Context, access, refresh string, accessExp, refreshExp time.Time) {
	accessCookie := &http.Cookie{
		Name:    AccessToken,
		Value:   access,
		Expires: accessExp,
		Path:    DefaultPath,
	}
	c.SetCookie(accessCookie)

	refreshCookie := &http.Cookie{
		Name:    RefreshToken,
		Value:   refresh,
		Expires: refreshExp,
		Path:    DefaultPath,
	}
	c.SetCookie(refreshCookie)
}
