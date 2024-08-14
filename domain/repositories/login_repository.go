package repositories

import (
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
)

type LoginRepository interface {
	Login(email, password string) (HubPlanner.LoginResponse, error)
	LoginGoogle(idToken string) (models.OAuthResponse, error)
}
