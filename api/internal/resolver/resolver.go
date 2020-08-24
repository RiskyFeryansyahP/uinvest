package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/awesomebusiness/uinvest/internal/service/authentication"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver serves as dependency injection
type Resolver struct {
	AuthenticationUC authentication.UsecaseAuthentication
}
