package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleUserInfo struct {
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

type GoogleTokenInfo struct {
	Issuer        string `json:"iss"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Audience      string `json:"aud"`
	ExpiresIn     int    `json:"exp"`
}

type OAuthResponse struct {
	Message   string `json:"message,omitempty"`
	Error     string `json:"error,omitempty"`
	Code      string `json:"code,omitempty"`
	Token     string `json:"token,omitempty"`
	UserEmail string `json:"userEmail,omitempty"`
}
