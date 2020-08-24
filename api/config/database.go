package config

import (
	"fmt"
	"os"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/facebookincubator/ent/dialect/sql"

	_ "github.com/lib/pq" // driver postgres
)

// Database is hold value client of ent
type Database struct {
	Client *ent.Client
}

// NewDatabase create new connection database
func NewDatabase() (*Database, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	drv, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed connecting database: %+v", err)
	}

	db := drv.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)

	client := ent.NewClient(ent.Driver(drv))

	return &Database{Client: client}, nil
}
