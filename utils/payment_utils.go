package utils

import (
	"LeaseEase/config"

	"github.com/omise/omise-go"
)

func NewOmiseClient() (*omise.Client, error) {
	publicKey := config.LoadConfig().OmisePublicKey
	secretKey := config.LoadConfig().OmiseSecretKey

	return omise.NewClient(publicKey, secretKey)
}
