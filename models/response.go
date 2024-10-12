package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"snaptalky/database"
	"strings"
	"time"
)

type Response struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `json:"userId"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Tone      string    `json:"tone" gorm:"not null"`
	Message   string    `json:"message" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	DeletedAt time.Time `json:"deletedAt" gorm:"autoDeleteTime"`
}

func (r *Response) BeforeCreate() error {
	validTones := map[string]bool{"flirting": true, "friendly": true, "formal": true}
	if r.Message == "" {
		return errors.New("message is required")
	}
	if r.Tone == "" {
		return errors.New("tone is required")
	}
	if _, isValid := validTones[r.Tone]; !isValid {
		return errors.New("tone must be oneof=flirting friendly formal")
	}
	return nil
}

// Add adds a new response, ensuring there are no more than 10 responses per tone for each user.
func (r *Response) Add() error {
	count := int64(0)
	if err := database.DB.Model(r).Where("user_id = ? AND tone = ?", r.UserID, r.Tone).Count(&count).Error; err != nil {
		return err
	}

	if count >= 10 {
		var oldestResponse Response
		if err := database.DB.Where("user_id = ? AND tone = ?", r.UserID, r.Tone).Order("created_at").First(&oldestResponse).Error; err != nil {
			return err
		}
		if err := database.DB.Delete(&oldestResponse).Error; err != nil {
			return err
		}
	}

	if err := database.DB.Create(r).Error; err != nil {
		return err
	}
	return nil
}

// GetMessagesByTone retrieves all responses for a given user and tone, and returns them
// as a concatenated string in the format "1-message. 2-message. 3-message"
func GetMessagesByTone(userID uuid.UUID, tone string) (string, error) {
	var responses []Response
	if err := database.DB.Where("user_id = ? AND tone = ?", userID, tone).Order("created_at").Find(&responses).Error; err != nil {
		return "", err
	}
	var messages []string
	for i, response := range responses {
		messages = append(messages, fmt.Sprintf("%d-%s", i+1, response.Message))
	}

	// Join the messages slice into a single string separated by ". "
	return strings.Join(messages, ". "), nil
}
