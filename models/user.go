package models

import (
	"time"
)

type Gender string

const (
	Male      Gender = "Male"
	Female    Gender = "Female"
	NonBinary Gender = "NonBinary" // Represents all other genders
)

type CommunicationStyle string

const (
	Normal  CommunicationStyle = "Default"
	Direct  CommunicationStyle = "Direct"
	Passive CommunicationStyle = "Passive"
)

type User struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	DeviceID           string             `json:"device_id" gorm:"uniqueIndex"`
	Age                int                `json:"age"`
	Gender             Gender             `json:"gender"`
	Language           string             `json:"language"`
	Bio                string             `json:"bio"`
	PublicID           string             `json:"public_id"`
	IsPremium          bool               `json:"is_premium"`
	LastScannedAt      time.Time          `json:"last_scanned_at"`
	ScanCount          int                `json:"scan_count"`
	CommunicationStyle CommunicationStyle `json:"communication_style"`
}
