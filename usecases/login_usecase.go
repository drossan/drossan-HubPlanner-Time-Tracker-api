package usecases

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/domain/repositories"
	"hubplanner-proxy-api/helpers"
)

type LoginUseCase struct {
	oauthRepository repositories.LoginRepository
}

func NewLoginUserUseCase(oauthRepository repositories.LoginRepository) *LoginUseCase {
	return &LoginUseCase{
		oauthRepository: oauthRepository,
	}
}

func (uc *LoginUseCase) Login(email, password string) (HubPlanner.LoginResponse, error) {
	return uc.oauthRepository.Login(email, password)
}

func (uc *LoginUseCase) LoginGoogle(idToken string) (models.OAuthResponse, error) {
	return uc.oauthRepository.LoginGoogle(idToken)
}

func (uc *LoginUseCase) RefreshAccessToken(refreshToken string) (HubPlanner.LoginResponse, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return HubPlanner.LoginResponse{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return HubPlanner.LoginResponse{}, jwt.ErrSignatureInvalid
	}

	userID := claims["user_id"].(string)

	newAccessToken, err := helpers.GenerateJWT(&HubPlanner.Resource{ID: userID})
	if err != nil {
		return HubPlanner.LoginResponse{}, err
	}

	return HubPlanner.LoginResponse{
		Token:        newAccessToken,
		RefreshToken: refreshToken,
	}, nil
}
