package models

import "time"

type TempUser struct {
	User     *User
	ExpireAt time.Time
}

type OTP struct {
	Email    string
	OTP      string
	ExpireAt time.Time
}
