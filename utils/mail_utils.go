package utils

import (
	"LeaseEase/config"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Generate a 6-digit OTP
func GenerateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n)
}

// Send OTP via email
func SendOTP(email, otp string) error {
	cfg := config.LoadConfig()

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailUser)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "LeaseEase - Your One-Time Password (OTP)")

	// HTML email body with company branding
	emailBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; }
				.header { text-align: center; font-size: 24px; font-weight: bold; color: #333; }
				.otp { font-size: 22px; font-weight: bold; color: #007BFF; }
				.footer { font-size: 14px; color: #555; text-align: center; margin-top: 20px; }
			</style>
		</head>
		<body>
			<div class="container">
				<p class="header">LeaseEase - OTP Verification</p>
				<p>Dear User,</p>
				<p>Your One-Time Password (OTP) for registration is:</p>
				<p class="otp">%s</p>
				<p>Please enter this code within the next 3 minutes to complete your registration.</p>
				<p>If you did not request this code, please ignore this email.</p>
				<p class="footer">Thank you for using LeaseEase!<br>Best regards,<br><strong>LeaseEase Team</strong></p>
			</div>
		</body>
		</html>
	`, otp)

	m.SetBody("text/html", emailBody)

	// Parse email port from config
	port, err := strconv.Atoi(cfg.EmailPort)
	if err != nil {
		return err
	}

	// Set up SMTP dialer
	d := gomail.NewDialer(cfg.EmailHost, port, cfg.EmailUser, cfg.EmailPassword)

	// Send email
	return d.DialAndSend(m)
}
