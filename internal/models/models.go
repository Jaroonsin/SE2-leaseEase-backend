package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Role     string // "lessor" or "lessee"
}

type Property struct {
	gorm.Model
	LessorID     uint
	Location     string
	Size         int
	Type         string
	Pricing      float64
	Availability string
}

type Request struct {
	gorm.Model
	PropertyID uint
	LesseeID   uint
	Status     string // "pending", "accepted", "declined"
}
