package models

import "time"

// User struct with proper GORM tags for relations and validations
// User struct with proper GORM tags for relations and validations
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Address    string `gorm:"size:255"`
	Birthday   time.Time
	Email      string     `gorm:"uniqueIndex;size:100;not null"`
	Password   string     `gorm:"size:100;not null"`
	UserType   string     `gorm:"size:50;not null"`
	Role       string     `gorm:"size:10;not null"` // lessor, lessee
	Properties []Property `gorm:"foreignKey:LessorID"`
}

// Property struct with relationships and mapping
type Property struct {
	PropertyID         uint             `gorm:"primaryKey"`
	LessorID           uint             `gorm:"not null"`
	Location           string           `gorm:"size:255;not null"`
	Size               string           `gorm:"size:50;not null"`
	Price              float64          `gorm:"not null"`
	AvailabilityStatus string           `gorm:"size:50;not null"`
	Lessor             User             `gorm:"foreignKey:LessorID"`
	PropertyReviews    []PropertyReview `gorm:"foreignKey:PropertyID"`
}

// Request struct with properly mapped fields
type Request struct {
	ID                     uint   `gorm:"primaryKey"`
	Purpose                string `gorm:"size:255"`
	ProposedMessage        string `gorm:"type:text"`
	Question               string `gorm:"type:text"`
	CreateAt               time.Time
	InterestedMarketSlotID uint
	LesseeID               uint `gorm:"not null"`
	Lessee                 User `gorm:"foreignKey:LesseeID;references:ID"`
}

// Review struct for reusable review fields
type Review struct {
	ReviewID      uint      `gorm:"primaryKey"`
	ReviewMessage string    `gorm:"type:text;not null"`
	Rating        int       `gorm:"not null"`
	TimeStamp     time.Time `gorm:"autoCreateTime"`
}

// LessorReview struct linking reviews with lessors
// A composite foreign key ensures proper mapping between tables
type LessorReview struct {
	ReviewID uint   `gorm:"primaryKey"`
	LessorID uint   `gorm:"primaryKey"`
	LesseeID uint   `gorm:"not null"`
	Review   Review `gorm:"foreignKey:ReviewID"`
	Lessor   User   `gorm:"foreignKey:LessorID"`
	Lessee   User   `gorm:"foreignKey:LesseeID"`
}

// PropertyReview struct linking reviews with properties
type PropertyReview struct {
	ReviewID   uint     `gorm:"primaryKey"`
	LesseeID   uint     `gorm:"not null"`
	PropertyID uint     `gorm:"not null"`
	Review     Review   `gorm:"foreignKey:ReviewID"`
	Lessee     User     `gorm:"foreignKey:LesseeID"`
	Property   Property `gorm:"foreignKey:PropertyID"`
}

// type Customer struct {
// 	ID           uint   `gorm:"primaryKey"`
// 	CustomerType string `gorm:"size:50"`
// }

// type Admin struct {
// 	ID uint `gorm:"primaryKey"`
// }

// type Lessor struct {
// 	ID      uint `gorm:"primaryKey"`
// 	AdminID uint
// 	Admin   Admin `gorm:"foreignKey:AdminID"`
// }

// type Lessee struct {
// 	ID      uint `gorm:"primaryKey"`
// 	AdminID uint
// 	Admin   Admin `gorm:"foreignKey:AdminID"`
// }

// type PremiumLessor struct {
// 	ID         uint `gorm:"primaryKey"`
// 	ExpireDate time.Time
// 	StartDate  time.Time
// }

// type Advertisement struct {
// 	AdvertisementID  uint   `gorm:"primaryKey"`
// 	Title            string `gorm:"size:255"`
// 	Description      string `gorm:"type:text"`
// 	PublishStartDate time.Time
// 	PublishEndDate   time.Time
// 	BannerURL        string `gorm:"size:255"`
// 	Link             string `gorm:"size:255"`
// 	ClientName       string `gorm:"size:100"`
// }

// type Transaction struct {
// 	TransactionID   uint   `gorm:"primaryKey"`
// 	TransactionType string `gorm:"size:50"`
// 	Currency        string `gorm:"size:10"`
// 	PaymentMethod   string `gorm:"size:50"`
// 	AccountID       string `gorm:"size:100"`
// 	Status          string `gorm:"size:50"`
// 	Amount          float64
// 	DateAndTime     time.Time
// 	LessorID        uint
// 	Lessor          User `gorm:"foreignKey:LessorID"`
// }

// type Problem struct {
// 	ProblemID   uint   `gorm:"primaryKey"`
// 	Subject     string `gorm:"size:255"`
// 	Description string `gorm:"type:text"`
// 	Status      string `gorm:"size:50"`
// 	CreateAt    time.Time
// }

// type ProblemTag struct {
// 	ProblemID  uint   `gorm:"primaryKey"`
// 	ProblemTag string `gorm:"size:50"`
// }

// type Solve struct {
// 	AdminID   uint
// 	ProblemID uint
// 	Admin     Admin   `gorm:"foreignKey:AdminID"`
// 	Problem   Problem `gorm:"foreignKey:ProblemID"`
// }

// type ChatMessage struct {
// 	MessageID        uint   `gorm:"primaryKey"`
// 	Message          string `gorm:"type:text"`
// 	ImageURL         string `gorm:"size:255"`
// 	TimeStamp        time.Time
// 	MessageDirection string `gorm:"size:50"`
// 	LessorID         uint
// 	Lessor           Lessor `gorm:"foreignKey:LessorID"`
// }

// type Report struct {
// 	MessageID  uint
// 	ProblemID  uint
// 	CustomerID uint
// 	Message    ChatMessage `gorm:"foreignKey:MessageID"`
// 	Problem    Problem     `gorm:"foreignKey:ProblemID"`
// }
