package v1

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"github.com/labstack/echo/v4"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := h.getAccessCookie(c)
		if err != nil {
			log.Println(err.Error())
			return newResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, id)
		return next(c)
	}
}

func (h *Handler) getAccessCookie(c echo.Context) (string, error) {
	accessCookie, err := c.Cookie(AccessToken)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return "", errors.New("you don't have any cookie")
		}
		return "", err
	}
	return h.authManager.Parse(accessCookie.Value)
}

func (h *Handler) parseAuthHeader(c echo.Context) (string, error) {
	header := c.Request().Header.Get(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.authManager.Parse(headerParts[1])
}

func getUserID(c echo.Context) (types.BinaryUUID, error) {
	return getIdByContext(c, userCtx)
}

func getIdByContext(c echo.Context, context string) (types.BinaryUUID, error) {
	idFromCtx := c.Get(context)

	idStr, ok := idFromCtx.(string)
	if !ok {
		return types.BinaryUUID{}, errors.New("userCtx is of invalid type")
	}

	id := types.ParseUUID(idStr)

	return id, nil
}
