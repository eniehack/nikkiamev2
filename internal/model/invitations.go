package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)


type Invitations struct {
    ID ulid.ULID `gorm:"type:char(26);primaryKey;uniqueIndex"`
	AvailableTime int `gorm:"default:1"`
	IsExpired bool `gorm:"default:false"`
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
