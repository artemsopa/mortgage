package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/artomsopun/mortgage/mortgage-api/internal/service"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/auth"
)

type Handler struct {
	services    *service.Services
	authManager auth.AuthManager
}

func NewHandler(services *service.Services, authManager auth.AuthManager) *Handler {
	return &Handler{
		services:    services,
		authManager: authManager,
	}
}

func (h *Handler) Init(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
	}
}
