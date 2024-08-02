package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
)

func GenerateJWT(user *HubPlanner.Resource) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    "Intranet API - Generate by IslaIT",
	}

	claims := models.Claim{
		UserID:           user.ID,
		Username:         user.Email,
		RegisteredClaims: registeredClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
