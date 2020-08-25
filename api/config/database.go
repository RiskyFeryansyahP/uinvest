package config

import (
	"github.com/awesomebusiness/uinvest/ent"

	_ "github.com/lib/pq" // driver postgres
)

// Database is hold value client of ent
type Database struct {
	Client *ent.Client
}
