package delivery

import (
	"net/http"

	"github.com/artomsopun/mortgage/mortgage-api/internal/config"
	v1 "github.com/artomsopun/mortgage/mortgage-api/internal/delivery/http/v1"
	"github.com/artomsopun/mortgage/mortgage-api/internal/service"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func (h *Handler) Init(cfg *config.Config) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      cfg.Limiter.RPS,
			Burst:     cfg.Limiter.Burst,
			ExpiresIn: cfg.Limiter.TTL,
		})))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	// Routes
	e.GET("/", healthCheck)
	h.initAPI(e)

	return e
}

func (h *Handler) initAPI(e *echo.Echo) {
	handlerV1 := v1.NewHandler(h.services, h.authManager)
	api := e.Group("/api")
	{
		handlerV1.Init(api)
	}
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Server is up and running",
	})
}
