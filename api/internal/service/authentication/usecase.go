package authentication

import (
	"context"

	"github.com/awesomebusiness/uinvest/internal/model"
)

// UsecaseAuthentication is interface that wraps authentication usecase method
type UsecaseAuthentication interface {
	AuthenticationValidation(ctx context.Context, input model.LoginInput) (*model.User, error)
	RegisterValidation(ctx context.Context, input model.RegisterInput) (*model.User, error)
}
