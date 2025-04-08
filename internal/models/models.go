package models

import "time"

// User struct with proper GORM tags for relations and validations
type User struct {
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"size:100;not null"`
	Password    string `gorm:"size:100;not null"`
	Name        string `gorm:"size:100;not null"`
	Address     string `gorm:"size:255"`
	ImageURL    string `gorm:"size:255"`
	CreatedAt   time.Time
	UserType    string `gorm:"size:50;not null"` // lessor, lessee
	ResetToken  string
	TokenExpiry time.Time
	Payments    []Payment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // One-to-Many Relationship
	Messages    []Message `gorm:"foreignKey:SenderID" json:"messages,omitempty"`
}

type Property struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string  `gorm:"size:100;not null"` // name of the property
	LessorID           uint    `gorm:"not null"`
	Location           string  `gorm:"size:255;not null"`
	Size               float64 `gorm:"size:50;not null"`
	AvgRating          float64 `gorm:"default:0"`
	ReviewCount        int     `gorm:"default:0"`
	Price              float64 `gorm:"not null"`
	AvailabilityStatus string  `gorm:"size:50;not null"`
	Details            string  `gorm:"type:text;not null"`
	Lessor             User    `gorm:"foreignKey:LessorID;references:ID"`
	ImageURL           string  `gorm:"size:255"`
}

// Request struct with properly mapped fields
type Reservation struct {
	ID                 uint   `gorm:"primaryKey"`
	Purpose            string `gorm:"size:255"`
	ProposedMessage    string `gorm:"type:text"`
	Question           string `gorm:"type:text"`
	Status             string `gorm:"size:50"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	InterestedProperty uint     `gorm:"not null"`
	Property           Property `gorm:"foreignKey:InterestedProperty;references:ID"`
	LesseeID           uint     `gorm:"not null"`
	Lessee             User     `gorm:"foreignKey:LesseeID;references:ID"`
}

// Review struct for reusable review fields
type Review struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	ReviewMessage string    `gorm:"type:text;not null"`
	Rating        int       `gorm:"not null"`
	TimeStamp     time.Time `gorm:"autoCreateTime"`
}

// LessorReview struct linking reviews with lessors
type LessorReview struct {
	ReviewID uint   `gorm:"primaryKey"`
	LessorID uint   `gorm:"not null"`
	LesseeID uint   `gorm:"not null"`
	Review   Review `gorm:"foreignKey:ReviewID;references:ID"`
	Lessor   User   `gorm:"foreignKey:LessorID;references:ID"`
	Lessee   User   `gorm:"foreignKey:LesseeID;references:ID"`
}

// PropertyReview struct linking reviews with properties
type PropertyReview struct {
	ReviewID   uint     `gorm:"primaryKey"`
	LesseeID   uint     `gorm:"not null"`
	PropertyID uint     `gorm:"not null"`
	Review     Review   `gorm:"foreignKey:ReviewID;references:ID"`
	Lessee     User     `gorm:"foreignKey:LesseeID;references:ID"`
	Property   Property `gorm:"foreignKey:PropertyID;references:ID"`
}
