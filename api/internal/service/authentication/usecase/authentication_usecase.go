package usecase

import (
	"context"
	"errors"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication"
)

// AuthenticationUsecase is usecase that handler authentication logic
type AuthenticationUsecase struct {
	authrepo authentication.RepositoryAuthentication
}

// NewAuthenticationUsecase create new usecase and injection to repository authentication
func NewAuthenticationUsecase(authrepo authentication.RepositoryAuthentication) authentication.UsecaseAuthentication {
	return &AuthenticationUsecase{
		authrepo: authrepo,
	}
}

// RegisterValidation validating register user
func (au *AuthenticationUsecase) RegisterValidation(ctx context.Context, input model.RegisterInput) (*ent.User, error) {
	if input.Firstname == "" {
		return nil, errors.New("firstname should not be empty")
	}

	if input.Lastname == "" {
		return nil, errors.New("lastname should not be empty")
	}

	if input.Email == "" {
		return nil, errors.New("email is not valid")
	}

	if input.Phonenumber == "" || len(input.Phonenumber) < 12 {
		return nil, errors.New("phone number is not valid")
	}

	if input.Password == "" || len(input.Password) < 8 {
		return nil, errors.New("password should not be empty or less than 8")
	}

	user, err := au.authrepo.CreateDataUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticationValidation validating authentication user
func (au *AuthenticationUsecase) AuthenticationValidation(ctx context.Context, input model.LoginInput) (*ent.User, error) {
	if input.Email == "" {
		return nil, errors.New("email is not valid")
	}

	if input.Password == "" || len(input.Password) < 8 {
		return nil, errors.New("password should not be empty or less than 8")
	}

	newUser, err := au.authrepo.GetDataUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
