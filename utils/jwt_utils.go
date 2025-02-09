package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT generates a JWT token for a user with a given user ID
func GenerateJWT(userID uint) (string, error) {
	// Define the secret key (should come from environment variables for security)
	secretKey := "your-secret-key" // Replace with config.LoadConfig().JWTSecret in your app

	// Define token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 3).Unix(), // Token expiration: 24 hours
		"iat":     time.Now().Unix(),                     // Issued at time
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

// ParseJWT parses and validates a JWT token, returning the claims if valid
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// Define the secret key
	secretKey := "your-secret-key" // Replace with config.LoadConfig().JWTSecret in your app

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract and return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
