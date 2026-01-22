package jwt

import (
	"fmt"
	"time"

	"github.com/ardianilyas/go-ticketing/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTService handles JWT operations
type JWTService struct {
	secretKey string
	expiresIn time.Duration
}

// NewJWTService creates a new JWT service
func NewJWTService() *JWTService {
	secret := config.Get("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is required in environment")
	}

	expiresInStr := config.Get("JWT_EXPIRES_IN")
	if expiresInStr == "" {
		expiresInStr = "24h"
	}

	expiresIn, err := time.ParseDuration(expiresInStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid JWT_EXPIRES_IN format: %v", err))
	}

	return &JWTService{
		secretKey: secret,
		expiresIn: expiresIn,
	}
}

// GenerateToken generates a new JWT token for a user
func (s *JWTService) GenerateToken(userID uuid.UUID, email string) (string, error) {
	claims := NewClaims(userID, email, s.expiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method is HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// GetExpiresIn returns the token expiration duration
func (s *JWTService) GetExpiresIn() time.Duration {
	return s.expiresIn
}
