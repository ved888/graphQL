package model

import (
	"github.com/google/uuid"
	"time"
)

type UserSession struct {
	UserID     uuid.UUID `json:"userId"`
	SessionKey string    `json:"sessionKey"`
	ExpiryAt   time.Time `json:"expiryAt"`
}

type Image struct {
	BucketName string `json:"bucket_name"`
	Path       string `json:"path"`
	UserId     string `json:"userId"`
}
