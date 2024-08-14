package usecases

import (
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/domain/repositories"
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
