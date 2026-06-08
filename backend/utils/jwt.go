package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims defines what data we store inside a JWT token.
// UserID and Role are our custom fields.
// jwt.RegisteredClaims handles standard fields like ExpiresAt, IssuedAt.
type Claims struct {
	UserID int    `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a signed JWT access token with userId and role.
// Expires in 15 minutes — short lived so it's safer.
func GenerateAccessToken(userId int, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token using HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateRefreshToken creates a signed JWT refresh token for the given userId.
// Expires in 7 days — longer lived, stored in Redis.
func GenerateRefreshToken(userId int) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken parses and validates a JWT token string.
// Returns the claims if valid, or an error if expired or tampered with.
func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HMAC — reject anything else
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
