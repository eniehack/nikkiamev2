package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type PostPassphrases struct {
	PostID ulid.ULID `gorm:"primaryKey;<-:create;type:char(26);uniqueIndex;"`
	Post Posts `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Passphrase string `gorm:"type:char(97)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
