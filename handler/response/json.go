package response

import (
	"github.com/labstack/echo/v4"
)

type ResponseMedia struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}
