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
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		Issuer:    "Hub Planner Proxy API - Generate by Secuoyas Experience",
	}

	claims := models.Claim{
		UserID:           user.ID,
		Username:         user.Email,
		RegisteredClaims: registeredClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}

func GetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrSignatureInvalid
	}
}
