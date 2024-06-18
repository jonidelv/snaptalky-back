package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"snaptalky/database"
	"snaptalky/utils"
	"time"
)

func init() {
	// AutoMigrate the User model
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		utils.LogError(err, "Failed to migrate database")
		return
	}
}

type gender string

const (
	male      gender = "male"
	female    gender = "female"
	nonBinary gender = "nonBinary" // Represents all other genders
)

type communicationStyle string

const (
	normal  communicationStyle = "normal"
	direct  communicationStyle = "direct"
	passive communicationStyle = "passive"
)

type tone string

const (
	flirting     tone = "flirting"
	friendly     tone = "friendly"
	professional tone = "professional"
)

type User struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	DeviceID           string             `json:"device_id" gorm:"uniqueIndex"`
	Age                int                `json:"age"`
	Gender             gender             `json:"gender"`
	Bio                string             `json:"bio"`
	PublicID           string             `json:"public_id" gorm:"uniqueIndex"`
	IsPremium          bool               `json:"is_premium" gorm:"default:false"`
	LastScannedAt      time.Time          `json:"last_scanned_at"`
	ScanCount          int                `json:"scan_count" gorm:"default:0"`
	CommunicationStyle communicationStyle `json:"communication_style" gorm:"default:normal"`
	Tone               tone               `json:"tone" gorm:"default:friendly"`
	UpdatedAt          time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt          time.Time          `json:"created_at" gorm:"autoCreateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.DeviceID == "" {
		return errors.New("device_id is required")
	}
	// Generate a unique PublicID
	u.PublicID = uuid.New().String()

	return nil
}

// IncrementScanCount UserIncrementScanCount IncrementScanCount atomically increments the ScanCount for the user.
func (u *User) IncrementScanCount() error {
	return database.DB.Model(u).Where("id = ?", u.ID).UpdateColumn(
		"scan_count", gorm.Expr("scan_count + ?", 1),
	).Error
}
