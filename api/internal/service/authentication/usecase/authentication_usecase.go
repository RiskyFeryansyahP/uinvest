package usecase

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication"
	"github.com/awesomebusiness/uinvest/util"
	"golang.org/x/crypto/scrypt"
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

	if !util.IsEmailValid(input.Email) {
		return nil, errors.New("email is invalid")
	}

	if !util.IsPhoneNumberValid(input.Phonenumber) {
		return nil, errors.New("phone number is not valid")
	}

	if input.Password == "" || len(input.Password) < 8 {
		return nil, errors.New("password should not be empty or less than 8")
	}

	// encrypt password
	encryptedPassword := hashPassword(input.Password)

	// update password with encrypted data
	input.Password = encryptedPassword

	user, err := au.authrepo.CreateDataUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticationValidation validating authentication user
func (au *AuthenticationUsecase) AuthenticationValidation(ctx context.Context, input model.LoginInput) (*ent.User, error) {
	if !util.IsEmailValid(input.Email) {
		return nil, errors.New("email is not valid")
	}

	if input.Password == "" || len(input.Password) < 8 {
		return nil, errors.New("password should not be empty or less than 8")
	}

	// encrypt password to get real string value of password
	encryptedPassword := hashPassword(input.Password)

	// update password with encrypted data
	input.Password = encryptedPassword

	newUser, err := au.authrepo.GetDataUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func hashPassword(password string) string {
	salt := []byte{0xf4, 0xc3, 0x49, 0x38, 0x00, 0xc5, 0x0f, 0xfc}

	encrypted, _ := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)

	return base64.StdEncoding.EncodeToString(encrypted)
}
