package model

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Posts struct {
	ID ulid.ULID `gorm:"primaryKey;<-:create;type:char(26);uniqueIndex;"`
	Title string `gorm:"type:varchar(20);"`
	AuthorID ulid.ULID `gorm:"uniqueIndex"`
	Author Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AuthorID;"`
	IsDraft bool `gorm:"default:false;"`
	Scope uint `gorm:"type:unsigned;"`
	content string `gorm:"type:text;"`
	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt
}
