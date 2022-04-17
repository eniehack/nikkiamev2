package main

import (
	"os"

	"github.com/eniehack/nikkiamev2/internal/config"
	"github.com/eniehack/nikkiamev2/internal/migrations"
	"github.com/urfave/cli/v2"
	"errors"
	"fmt"
)

func main() {
	var configFilePath string
	app := cli.NewApp()
	app.Name = "migrator"
	app.Usage = "migrates database"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config, c",
			Value:       "../../config.toml",
			Usage:       "specify config file",
			Destination: &configFilePath,
			EnvVars:      []string{"NIKKIAME_CONFIG"},
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if ctx.String("config") == "" {
			return errors.New("config is required")
		}

		configData, err := config.LoadConfigFile(ctx.String("config"))
		if err != nil {
			return fmt.Errorf("failed to connect database: %s", err)
		}

		db, err := config.SetupDatabase(configData)
		if err != nil {
			return fmt.Errorf("failed to connect database: %s", err)
		}

		if err := migrations.Migrations(db); err != nil {
			return fmt.Errorf("failed to migrate: %s", err)
		}
		return nil
	}
	app.Run(os.Args)
}
