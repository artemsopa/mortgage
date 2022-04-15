package v1

import "github.com/labstack/echo/v4"

type response struct {
	Message string `json:"message"`
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

func newResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, response{message})
}
