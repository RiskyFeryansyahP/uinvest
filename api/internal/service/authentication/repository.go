package authentication

import (
	"context"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
)

// RepositoryAuthentication is interface that wraps the authentication repository method.
type RepositoryAuthentication interface {
	GetDataUser(ctx context.Context, input model.LoginInput) (*ent.User, error)
	CreateDataUser(ctx context.Context, input model.RegisterInput) (*ent.User, error)
}
