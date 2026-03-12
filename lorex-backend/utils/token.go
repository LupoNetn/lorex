package utils

import (
	"errors"
	"time"
	"strings"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Secret key for JWT signing - ideally this comes from environment variables
var jwtSecret = []byte("super-secret-key-change-in-production")

// CustomClaims represents the JWT claims
type CustomClaims struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	CompanyID string `json:"company_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateTokens creates both an access token and a refresh token
func GenerateTokens(companyID uuid.UUID, email string, id uuid.UUID) (string, string, error) {
	// 1. Generate Access Token (short lived)
	accessTokenClaims := CustomClaims{
		ID:        id.String(),
		Role:      "company",
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
		ID:        id.String(),
		Role:      "company",
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

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("unexpected signing method", "error", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		slog.Error("Invalid token", "error", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	slog.Error("Invalid token")
	return nil, errors.New("invalid token")
}


func ExtractToken(tokenString string) string {
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}
