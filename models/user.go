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
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

type CommunicationStyle string

// CommunicationStyle enum values
const (
	CommunicationStyleDefault CommunicationStyle = "default"
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
	DeviceID           string             `json:"deviceID" gorm:"uniqueIndex"`
	Platform           string             `json:"platform,omitempty"`
	Age                int                `json:"age,omitempty"`
	Gender             Gender             `json:"gender,omitempty"`
	Bio                string             `json:"bio,omitempty"`
	PublicID           string             `json:"publicID" gorm:"uniqueIndex"`
	IsPremium          bool               `json:"isPremium" gorm:"default:false"`
	IsPremiumAt        time.Time          `json:"IsPremiumAt,omitempty"`
	LastScannedAt      time.Time          `json:"lastScannedAt,omitempty"`
	ScanCount          int                `json:"scanCount" gorm:"default:0"`
	CommunicationStyle CommunicationStyle `json:"communicationStyle" gorm:"default:default"`
	UpdatedAt          time.Time          `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedAt          time.Time          `json:"createdAt" gorm:"autoCreateTime"`
	DeletedAt          time.Time          `json:"deletedAt" gorm:"autoDeleteTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.DeviceID == "" {
		return errors.New("deviceID is required")
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
