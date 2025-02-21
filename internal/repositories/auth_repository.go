package repositories

import (
	"LeaseEase/internal/models"
	"LeaseEase/utils"
	"sync"
	"time"

	"gorm.io/gorm"
)

type authRepository struct {
	db        *gorm.DB
	userStore map[string]models.TempUser
	otpStore  map[string]models.OTP
	mu        sync.Mutex
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	repo := &authRepository{
		db:        db,
		userStore: make(map[string]models.TempUser),
		otpStore:  make(map[string]models.OTP),
	}
	go repo.cleanupRoutine()
	return repo
}

func (r *authRepository) FindEmailExisted(email string) bool {
	var user models.User
	r.db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Save บันทึกบัญชีชั่วคราว
func (r *authRepository) SaveTempUser(user models.TempUser) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userStore[user.User.Email] = user
}

// FindByEmail ค้นหาบัญชีชั่วคราว
func (r *authRepository) FindByEmailTempUser(email string) (models.TempUser, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.userStore[email]
	return user, exists
}

// Delete ลบบัญชีชั่วคราว
func (r *authRepository) DeleteTempUser(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.userStore, email)
}

// Save บันทึก OTP
func (r *authRepository) SaveOTP(otp models.OTP) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.otpStore[otp.Email] = otp
}

// FindByEmail ค้นหา OTP จากอีเมล
func (r *authRepository) FindByEmailOTP(email string) (models.OTP, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	otp, exists := r.otpStore[email]
	return otp, exists
}

// Delete ลบ OTP หลังจากใช้
func (r *authRepository) DeleteOTP(email string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.otpStore, email)
}

// CleanupExpiredOTP ลบ OTP ที่หมดอายุ
func (r *authRepository) cleanupExpiredOTP() {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()
	for email, otp := range r.otpStore {
		if otp.ExpireAt.Before(now) {
			delete(r.otpStore, email)
		}
	}
}

// CleanupExpiredUsers ลบข้อมูลที่หมดอายุ
func (r *authRepository) cleanupExpiredUsers() {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()
	for email, user := range r.userStore {
		if user.ExpireAt.Before(now) {
			delete(r.userStore, email)
		}
	}
}

// cleanupRoutine Goroutine สำหรับล้าง OTP ทุก 10 วินาที
func (r *authRepository) cleanupRoutine() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		r.cleanupExpiredOTP()
		r.cleanupExpiredUsers()
	}
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) SaveResetToken(user *models.User, token string, expiry time.Time) error {
	hashedToken := utils.HashToken(token)
	user.ResetToken = hashedToken
	user.TokenExpiry = expiry
	return r.db.Save(user).Error
}

func (r *authRepository) UpdateUserPassword(user *models.User, hashedPassword string) error {
	user.Password = hashedPassword
	user.ResetToken = "" // Clear reset token
	return r.db.Save(user).Error
}
