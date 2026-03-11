package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Secret key for JWT signing - ideally this comes from environment variables
var jwtSecret = []byte("super-secret-key-change-in-production")

// CustomClaims represents the JWT claims
type CustomClaims struct {
	CompanyID string `json:"company_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateTokens creates both an access token and a refresh token
func GenerateTokens(companyID uuid.UUID, email string) (string, string, error) {
	// 1. Generate Access Token (short lived)
	accessTokenClaims := CustomClaims{
		CompanyID: companyID.String(),
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // 15 mins
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   companyID.String(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAccessToken, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// 2. Generate Refresh Token (long lived)
	refreshTokenClaims := CustomClaims{
		CompanyID: companyID.String(),
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   companyID.String(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}
