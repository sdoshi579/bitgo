package models

import (
	"github.com/google/uuid"
	"github.com/sodhi579/bitgo/app/notification/value_objects"
)

type CreateNotificationRequest struct {
	UserID            uuid.UUID
	CurrentPrice      float64
	Volume            float64
	IntraDayHighPrice float64
	MarketCap         float64
	Status            value_objects.Status
	IsDeleted         bool
}
