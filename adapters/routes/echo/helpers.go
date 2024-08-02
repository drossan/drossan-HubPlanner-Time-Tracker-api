package echo

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"hubplanner-proxy-api/domain/models"
)

func GetUserID(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.Claim)
	return claims.UserID
}
