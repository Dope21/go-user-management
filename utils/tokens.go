package utils

import (
	"fmt"
	"time"
	"user-management/constants/configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func tokenClaim(userID uuid.UUID, duration time.Duration) jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(duration).Unix(),
	}
}

func tokenSign(userID uuid.UUID, duration time.Duration, secretKey string) (string, error) {
	claim := tokenClaim(userID, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}

func GenerateTokens(userID uuid.UUID) (string, string, error) {
	cfg := configs.LoadConfig()

	accessToken, err := tokenSign(userID, 15*time.Minute, cfg.AppTokenKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	refreshToken, err := tokenSign(userID, 7*24*time.Hour, cfg.AppTokenKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return accessToken, refreshToken, nil
}