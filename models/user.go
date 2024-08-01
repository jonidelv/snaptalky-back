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
	ID                 uuid.UUID          `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
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
	DeletedAt          time.Time          `json:"deleted_at" gorm:"autoDeleteTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.DeviceID == "" {
		return errors.New("device_id is required")
	}
	// Generate a unique PublicID
	u.PublicID = uuid.New().String()

	return nil
}

// IncrementScanCount increments the ScanCount field for the User instance.
// This operation is performed atomically to ensure thread safety.
// It updates the scan_count column in the database by incrementing its value by 1.
func (u *User) IncrementScanCount() error {
	return database.DB.Model(u).Where("id = ?", u.ID).UpdateColumn(
		"scan_count", gorm.Expr("scan_count + ?", 1),
	).Error
}
