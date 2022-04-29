package model

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Users struct {
	Ulid string `gorm:"primaryKey;<-:create;type:char(26);uniqueIndex"`
	UserId string `gorm:"unique;type:varchar(15);uniqueIndex;not null"`
	Name string `gorm:"size:20;not null"`
	Password string `gorm:"not null;type:char(97)"`
	IsAdmin bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *Users) BeforeCreate(db *gorm.DB) error {
	if u.Ulid == "" {
		ulid, err := ulid.New(ulid.Now(), rand.Reader)
		if err != nil {
			return err
		}
		u.Ulid = ulid.String()
	}
	return nil
}
