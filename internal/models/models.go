package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100"`
	Address  string `gorm:"size:255"`
	Birthday time.Time
	Email    string `gorm:"unique;size:100"`
	Password string `gorm:"size:100"`
	UserType string `gorm:"size:50"`
}

type Customer struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerType string `gorm:"size:50"`
}

type Admin struct {
	ID uint `gorm:"primaryKey"`
}

type Lessor struct {
	ID      uint `gorm:"primaryKey"`
	AdminID uint
	Admin   Admin `gorm:"foreignKey:AdminID"`
}

type Lessee struct {
	ID      uint `gorm:"primaryKey"`
	AdminID uint
	Admin   Admin `gorm:"foreignKey:AdminID"`
}

type PremiumLessor struct {
	ID         uint `gorm:"primaryKey"`
	ExpireDate time.Time
	StartDate  time.Time
}

type Advertisement struct {
	ID               uint   `gorm:"primaryKey"`
	Title            string `gorm:"size:255"`
	Description      string `gorm:"type:text"`
	PublishStartDate time.Time
	PublishEndDate   time.Time
	BannerURL        string `gorm:"size:255"`
	Link             string `gorm:"size:255"`
	ClientName       string `gorm:"size:100"`
	AdminID          uint
	Admin            Admin `gorm:"foreignKey:AdminID"`
}

type MarketSlot struct {
	MarketSlotID       uint `gorm:"primaryKey"`
	LessorID           uint
	Location           string `gorm:"size:255"`
	Size               string `gorm:"size:50"`
	Price              float64
	AvailabilityStatus string `gorm:"size:50"`
	AdminID            uint
	Admin              Admin `gorm:"foreignKey:AdminID"`
}

type Request struct {
	ID                     uint   `gorm:"primaryKey"`
	Purpose                string `gorm:"size:255"`
	ProposedMessage        string `gorm:"type:text"`
	Question               string `gorm:"type:text"`
	CreateAt               time.Time
	InterestedMarketSlotID uint
	LesseeID               uint
	Lessee                 Lessee `gorm:"foreignKey:LesseeID"`
}

type Transaction struct {
	TransactionID   uint   `gorm:"primaryKey"`
	TransactionType string `gorm:"size:50"`
	Currency        string `gorm:"size:10"`
	PaymentMethod   string `gorm:"size:50"`
	AccountID       string `gorm:"size:100"`
	Status          string `gorm:"size:50"`
	Amount          float64
	DateAndTime     time.Time
	LessorID        uint
	Lessor          Lessor `gorm:"foreignKey:LessorID"`
}

type Review struct {
	ReviewID      uint   `gorm:"primaryKey"`
	ReviewMessage string `gorm:"type:text"`
	Rating        int
	TimeStamp     time.Time
}

type LessorReview struct {
	ReviewID uint
	LessorID uint
	LesseeID uint
	Review   Review `gorm:"foreignKey:ReviewID"`
}

type SlotReview struct {
	ReviewID     uint
	LesseeID     uint
	MarketSlotID uint
	Review       Review `gorm:"foreignKey:ReviewID"`
}

type Problem struct {
	ProblemID   uint   `gorm:"primaryKey"`
	Subject     string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"size:50"`
	CreateAt    time.Time
}

type ProblemTag struct {
	ProblemID  uint   `gorm:"primaryKey"`
	ProblemTag string `gorm:"size:50"`
}

type Solve struct {
	AdminID   uint
	ProblemID uint
	Admin     Admin   `gorm:"foreignKey:AdminID"`
	Problem   Problem `gorm:"foreignKey:ProblemID"`
}

type ChatMessage struct {
	MessageID        uint   `gorm:"primaryKey"`
	Message          string `gorm:"type:text"`
	ImageURL         string `gorm:"size:255"`
	TimeStamp        time.Time
	MessageDirection string `gorm:"size:50"`
	LessorID         uint
	Lessor           Lessor `gorm:"foreignKey:LessorID"`
}

type Report struct {
	MessageID  uint
	ProblemID  uint
	CustomerID uint
	Message    ChatMessage `gorm:"foreignKey:MessageID"`
	Problem    Problem     `gorm:"foreignKey:ProblemID"`
}
