package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"snaptalky/database"
	"time"
)

type User struct {
	ID                 uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DeviceID           string     `json:"deviceID" gorm:"uniqueIndex;not null"`
	Platform           string     `json:"platform,omitempty" gorm:"not null"`
	Age                int        `json:"age,omitempty"`
	Gender             string     `json:"gender,omitempty"`
	Bio                string     `json:"bio,omitempty"`
	PublicID           string     `json:"publicID" gorm:"uniqueIndex;not null"`
	IsPremium          bool       `json:"isPremium" gorm:"default:true"` // TODO change this to false
	IsPremiumAt        time.Time  `json:"IsPremiumAt,omitempty"`
	LastScannedAt      time.Time  `json:"lastScannedAt,omitempty"`
	ScanCount          int        `json:"scanCount" gorm:"default:0"`
	UsagesCount        int        `json:"usagesCount" gorm:"default:0"`
	CommunicationStyle string     `json:"communicationStyle" gorm:"default:default"`
	UpdatedAt          *time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedAt          *time.Time `json:"createdAt" gorm:"autoCreateTime"`
	DeletedAt          *time.Time `json:"deletedAt" gorm:"autoDeleteTime"`
}

func (u *User) validateGender() error {
	if u.Gender != "" && !(u.Gender == "male" || u.Gender == "female" || u.Gender == "other") {
		return errors.New("invalid gender value")
	}
	return nil
}

func (u *User) validateCommunicationStyle() error {
	if !(u.CommunicationStyle == "default" || u.CommunicationStyle == "direct" || u.CommunicationStyle == "passive") {
		return errors.New("invalid communication style value")
	}
	return nil
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.DeviceID == "" {
		return errors.New("deviceID is required")
	}
	u.PublicID = uuid.New().String()
	if err := u.validateGender(); err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	if err := u.validateGender(); err != nil {
		return err
	}
	if err := u.validateCommunicationStyle(); err != nil {
		return err
	}

	return nil
}

// IncrementCountsAndUsages increments the ScanCount and UsagesCount field for the User instance.
// This operation is performed atomically to ensure thread safety.
// It updates the scan_count column in the database by incrementing its value by 1 and the usages_count byt the usages
func (u *User) IncrementCountsAndUsages(usages int) error {
	return database.DB.Model(u).Where("id = ?", u.ID).Updates(map[string]interface{}{
		"scan_count":   gorm.Expr("scan_count + ?", 1),
		"usages_count": gorm.Expr("usages_count + ?", usages),
	}).Error
}
