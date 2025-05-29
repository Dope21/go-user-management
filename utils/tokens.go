package utils

import (
	"errors"
	"fmt"
	"time"
	"user-management/constants/configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

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

	refreshToken, err := tokenSign(userID, 7*24*time.Second, cfg.AppTokenKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return accessToken, refreshToken, nil
}

func VerifyToken(tokenStr string) (uuid.UUID, error) {
	cfg := configs.LoadConfig()

	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.AppTokenKey), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return claims.UserID, errors.New("token is expired")
	}

	return claims.UserID, nil

}