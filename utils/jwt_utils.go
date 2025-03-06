package utils

import (
	"errors"
	"time"

	"LeaseEase/config"
	"LeaseEase/internal/dtos"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT generates a JWT token for a user with a given user ID
func GenerateJWT(user dtos.JWTDTO) (string, error) {
	// Define the secret key (should come from environment variables for security)
	secretKey := config.LoadConfig().JWTApiSecret

	// Define token claims
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 3).Unix(), // Token expiration: 24 hours
		"iat":     time.Now().Unix(),                    // Issued at time
		"role":    user.Role,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := config.LoadConfig().JWTApiSecret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	} else {
		return nil, errors.New("invalid token: missing expiration")
	}

	return claims, nil
}
