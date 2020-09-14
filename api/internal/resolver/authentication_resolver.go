package resolver

import (
	"context"

	"github.com/awesomebusiness/uinvest/internal/model"
)

func (r *queryResolver) Login(ctx context.Context, input model.LoginInput) (*model.User, error) {
	result, err := r.AuthenticationUC.AuthenticationValidation(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *queryResolver) Register(ctx context.Context, input model.RegisterInput) (*model.User, error) {
	result, err := r.AuthenticationUC.RegisterValidation(ctx, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
