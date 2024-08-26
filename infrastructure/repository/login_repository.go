package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/helpers"
)

type oauthRepository struct{}

func NewLoginConnectionRepository() *oauthRepository {
	return &oauthRepository{}
}

func (r *oauthRepository) Login(email, password string) (HubPlanner.LoginResponse, error) {
	var result HubPlanner.LoginResponse

	// Preparar los datos para la solicitud
	data := map[string]string{
		"userName": email,
		"password": password,
	}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", os.Getenv("API_URI_COMPANY")+"/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Secuoyas Experiences - Time Tracking (daniel.rossello@secuoyas.com)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &result)

	if result.Status {
		resources, err := r.recoveryUser(email)
		if err != nil {
			return result, err
		}

		if len(resources) > 0 {
			accessToken, err := helpers.GenerateJWT(&resources[0])
			if err != nil {
				return result, err
			}

			refreshToken, err := helpers.GenerateRefreshToken(resources[0].ID)
			if err != nil {
				return result, err
			}

			result.Token = accessToken
			result.RefreshToken = refreshToken
		}
	}

	return result, nil
}

func (r *oauthRepository) LoginGoogle(idToken string) (models.OAuthResponse, error) {
	var result models.OAuthResponse

	userInfoEndpoint := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
	req, err := http.NewRequest("GET", userInfoEndpoint, nil)
	if err != nil {
		return result, err
	}

	req.Header.Set("Authorization", "Bearer "+idToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var userInfo models.GoogleUserInfo
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &userInfo)

	if userInfo.Email != "" && userInfo.VerifiedEmail {
		resources, err := r.recoveryUser(userInfo.Email)
		if err != nil {
			return result, err
		}

		if len(resources) > 0 {
			accessToken, err := helpers.GenerateJWT(&resources[0])
			if err != nil {
				return result, err
			}

			refreshToken, err := helpers.GenerateRefreshToken(resources[0].ID)
			if err != nil {
				return result, err
			}

			result.Token = accessToken
			result.RefreshToken = refreshToken
			result.UserEmail = userInfo.Email
		}
	}

	result.Error = "Usuario no válido"
	result.Message = "Usuario no válido"

	return result, nil
}

func (r *oauthRepository) recoveryUser(username string) ([]HubPlanner.Resource, error) {
	var resources []HubPlanner.Resource

	url := os.Getenv("API_URL") + "/resource/search"
	method := "POST"

	payload := fmt.Sprintf(`{
		"email": {
			"$in": ["%s"]
		}
	}`, username)

	bodyBytes, err := helpers.MakeHTTPRequest(
		method,
		url,
		os.Getenv("API_TOKEN"),
		"application/json",
		strings.NewReader(payload),
	)

	if err != nil {
		return resources, err
	}

	_ = json.Unmarshal(bodyBytes, &resources)
	return resources, nil
}
