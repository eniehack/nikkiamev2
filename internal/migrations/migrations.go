package migrations

import (
	"github.com/eniehack/nikkiamev2/internal/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) error {
	migrations := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				return db.Migrator().AutoMigrate(&model.Users{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("users")
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return db.Migrator().AutoMigrate(&model.Invitations{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("invitations")
			},
		},
		{
			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				return db.Migrator().AutoMigrate(&model.Posts{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("posts")
			},
		},
		{
			ID: "4",
			Migrate: func(tx *gorm.DB) error {
				if db.Migrator().HasTable("posts_passphrases") {
					return db.Migrator().RenameTable("posts_passphrases", "post_passphrases")
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if db.Migrator().HasTable("post_passphrases") {
					return db.Migrator().RenameTable("post_passphrases", "posts_passphrases")
				}
				return nil
			},
		},
		{
			ID: "5",
			Migrate: func(tx *gorm.DB) error {
				return db.Migrator().AutoMigrate(&model.PostPassphrases{})
			},
			Rollback: func(tx *gorm.DB) error {
				return db.Migrator().DropTable("post_passphrases")
			},
		},
	})

	return migrations.Migrate()
}
