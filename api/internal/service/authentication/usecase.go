package authentication

import (
	"context"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
)

// UsecaseAuthentication is interface that wraps authentication usecase method
type UsecaseAuthentication interface {
	AuthenticationValidation(ctx context.Context, input model.LoginInput) (*ent.User, error)
	RegisterValidation(ctx context.Context, input model.RegisterInput) (*ent.User, error)
}
