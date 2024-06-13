package models

import (
	"time"
)

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

type User struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	DeviceID           string             `json:"device_id" gorm:"uniqueIndex"`
	Age                int                `json:"age"`
	Gender             gender             `json:"gender"`
	Language           string             `json:"language"`
	Bio                string             `json:"bio"`
	PublicID           string             `json:"public_id"`
	IsPremium          bool               `json:"is_premium"`
	LastScannedAt      time.Time          `json:"last_scanned_at"`
	ScanCount          int                `json:"scan_count"`
	CommunicationStyle communicationStyle `json:"communication_style"`
	UpdatedAt          time.Time          `json:"updated_at"`
}
