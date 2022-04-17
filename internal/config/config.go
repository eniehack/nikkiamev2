package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DBTYPE_SQLITE   = "sqlite3"
	DBTYPE_POSTGRES = "postgres"
	DBTYPE_MYSQL    = "mysql"
)

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

type Config struct {
	DB DatabaseConfig `toml:"database"`
}

func LoadConfigFile(filepath string) (*Config, error) {
	var config = new(Config)
	if _, err := toml.DecodeFile(filepath, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func SetupDatabase(config *Config) (*gorm.DB, error) {
	switch config.DB.Driver {
	case DBTYPE_SQLITE:
		db, err := gorm.Open(sqlite.Open(config.DB.DSN), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("setupDatabase(): failed to open database: %s", err)
		}
		return db, nil
	default:
		return nil, errors.New("setupDatabase(): this driver is not supported")
	}
}
