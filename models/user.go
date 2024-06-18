package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"snaptalky/database"
	"time"
)

type Gender string

// Gender enum values
const (
	GenderMale      Gender = "male"
	GenderFemale    Gender = "female"
	GenderNonBinary Gender = "nonBinary"
)

type CommunicationStyle string

// CommunicationStyle enum values
const (
	CommunicationStyleNormal  CommunicationStyle = "normal"
	CommunicationStyleDirect  CommunicationStyle = "direct"
	CommunicationStylePassive CommunicationStyle = "passive"
)

type Tone string

// Tone enum values
const (
	ToneFlirting     Tone = "flirting"
	ToneFriendly     Tone = "friendly"
	ToneProfessional Tone = "professional"
	ToneCustom       Tone = "custom"
)

type User struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	DeviceID           string             `json:"device_id" gorm:"uniqueIndex"`
	Platform           string             `json:"platform,omitempty"`
	Age                int                `json:"age,omitempty"`
	Gender             Gender             `json:"gender,omitempty"`
	Bio                string             `json:"bio,omitempty"`
	PublicID           string             `json:"public_id" gorm:"uniqueIndex"`
	IsPremium          bool               `json:"is_premium" gorm:"default:false"`
	LastScannedAt      time.Time          `json:"last_scanned_at,omitempty"`
	ScanCount          int                `json:"scan_count" gorm:"default:0"`
	CommunicationStyle CommunicationStyle `json:"communication_style" gorm:"default:normal"`
	Tone               Tone               `json:"tone" gorm:"default:friendly"`
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
