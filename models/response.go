package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jonidelv/snaptalky-back/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Response struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID  `json:"userID"`
	User      User       `json:"user" gorm:"foreignKey:UserID"`
	Tone      string     `json:"tone" gorm:"not null"`
	Message   string     `json:"message" gorm:"not null"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"autoDeleteTime"`
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

// Add adds a new response, ensuring there are no duplicates and no more than 10 responses per tone for each user.
func (r *Response) Add() error {
	// Start a transaction to ensure atomicity of operations
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Check for existing Response with the same UserID, Tone, and Message
		var existing Response
		err := tx.Where("user_id = ? AND tone = ? AND message = ?", r.UserID, r.Tone, r.Message).First(&existing).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// An unexpected error occurred while querying
			return fmt.Errorf("error checking for existing response: %w", err)
		}
		if err == nil {
			return errors.New("duplicate response found. Skipping addition")
		}

		// 2. Count existing Responses for the UserID and Tone
		var count int64
		if err := tx.Model(&Response{}).
			Where("user_id = ? AND tone = ?", r.UserID, r.Tone).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error counting responses: %w", err)
		}

		// 3. If count >= 10, delete the oldest Response
		if count >= 10 {
			var oldestResponse Response
			if err := tx.
				Where("user_id = ? AND tone = ?", r.UserID, r.Tone).
				Order("created_at ASC").
				First(&oldestResponse).Error; err != nil {
				return fmt.Errorf("error fetching oldest response: %w", err)
			}
			if err := tx.Delete(&oldestResponse).Error; err != nil {
				return fmt.Errorf("error deleting oldest response: %w", err)
			}
		}

		// 4. Create the new Response
		if err := tx.Create(r).Error; err != nil {
			return fmt.Errorf("error creating new response: %w", err)
		}

		return nil
	})
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
