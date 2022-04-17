package model

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Users struct {
	Ulid ulid.ULID `gorm:"primaryKey;<-:create;type:char(26);uniqueIndex"`
	UserID string `gorm:"unique;type:varchar(15);uniqueIndex;not null"`
	Name string `gorm:"size:20;not null"`
	Password string `gorm:"not null;type:char(97)"`
	IsAdmin bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
