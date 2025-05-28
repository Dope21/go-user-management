package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte("secret-key")

func tokenClaim(userID uuid.UUID, duration time.Duration) jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(duration).Unix(),
	}
}

func tokenSign(userID uuid.UUID, duration time.Duration) (string, error) {
	claim := tokenClaim(userID, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secretKey)
}

func GenerateTokens(userID uuid.UUID) (string, string, error) {

	accessToken, err := tokenSign(userID, 15*time.Minute)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	refreshToken, err := tokenSign(userID, 7*24*time.Hour)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return accessToken, refreshToken, nil
}