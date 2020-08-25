package config

import (
	"fmt"
	"os"

	"github.com/awesomebusiness/uinvest/ent"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Config hold value all of configuration
type Config struct {
	Database *Database
}

// NewConfigMap create new configuration
func NewConfigMap() (*Config, error) {

	client, err := databaseConnection()
	if err != nil {
		return nil, err
	}

	database := &Database{
		Client: client,
	}

	return &Config{Database: database}, nil
}

func databaseConnection() (*ent.Client, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	drv, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed connecting database: %+v", err)
	}

	db := drv.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)

	client := ent.NewClient(ent.Driver(drv))

	return client, nil
}
