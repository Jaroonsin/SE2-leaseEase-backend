package utils

import (
	"LeaseEase/config"
	"LeaseEase/internal/dtos"
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

// SendPasswordResetEmail sends a password reset link to the user's email.
func SendPasswordResetEmail(req *dtos.ResetPassRequestDTO, resetURL string) error {
	cfg := config.LoadConfig()
	email := req.Email

	// Create email message
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailUser)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "LeaseEase - Password Reset Request")

	// HTML email body
	emailBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; }
				.header { text-align: center; font-size: 24px; font-weight: bold; color: #333; }
				.reset-link { display: block; width: 200px; margin: 20px auto; padding: 10px; text-align: center; 
				              background-color: #007BFF; color: white; text-decoration: none; border-radius: 5px; }
				.footer { font-size: 14px; color: #555; text-align: center; margin-top: 20px; }
			</style>
		</head>
		<body>
			<div class="container">
				<p class="header">LeaseEase - Password Reset</p>
				<p>Dear User,</p>
				<p>You requested a password reset. Click the link below to reset your password:</p>
				<a href="%s" class="reset-link">Reset Password</a>
				<p>If you did not request this, please ignore this email.</p>
				<p class="footer">Thank you for using LeaseEase!<br>Best regards,<br><strong>LeaseEase Team</strong></p>
			</div>
		</body>
		</html>
	`, resetURL)

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

func SendLessorAcceptanceEmail(req *dtos.AcceptReservationDTO) error {
	cfg := config.LoadConfig()
	lesseeEmail := req.LesseeEmail
	propertyName := req.PropertyName

	// Create email message
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailUser)
	m.SetHeader("To", lesseeEmail)
	m.SetHeader("Subject", "LeaseEase - Rental Request Approved")

	// HTML email body
	emailBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; }
				.header { text-align: center; font-size: 24px; font-weight: bold; color: #333; }
				.content { font-size: 16px; color: #444; }
				.footer { font-size: 14px; color: #555; text-align: center; margin-top: 20px; }
			</style>
		</head>
		<body>
			<div class="container">
				<p class="header">LeaseEase - Rental Request Approved</p>
				<p class="content">Dear Valued Lessee,</p>
				<p class="content">We are pleased to inform you that your request to lease <strong>%s</strong> has been formally approved.</p>
				<p class="content">Kindly log in to your account at your earliest convenience to review the terms and proceed with the necessary formalities.</p>
				<p class="content">Should you have any questions or require further assistance, please do not hesitate to contact us.</p>
				<p class="footer">Thank you for choosing LeaseEase.<br>Best regards,<br><strong>The LeaseEase Team</strong></p>
			</div>
		</body>
		</html>
	`, propertyName)

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
