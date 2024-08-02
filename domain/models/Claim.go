package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claim Token de user
type Claim struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
