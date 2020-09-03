package repository

import (
	"context"
	"fmt"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/ent/user"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication"
	"github.com/awesomebusiness/uinvest/pkg"

	log "github.com/sirupsen/logrus"
)

// AuthenticationRepository is repository that handle user authentication
type AuthenticationRepository struct {
	DB      *ent.Client
	Twillio pkg.TwillioMessage
}

// NewAuthenticationRepository create new repository with connected to database
func NewAuthenticationRepository(db *ent.Client, twillioClient pkg.TwillioMessage) authentication.RepositoryAuthentication {
	return &AuthenticationRepository{
		DB:      db,
		Twillio: twillioClient,
	}
}

// CreateDataUser create new user profile to database
func (ar *AuthenticationRepository) CreateDataUser(ctx context.Context, input model.RegisterInput) (*ent.User, error) {
	newUser, err := ar.DB.User.
		Create().
		SetFirstname(input.Firstname).
		SetLastname(input.Lastname).
		SetEmail(input.Email).
		SetPhonenumber(input.Phonenumber).
		SetPassword(input.Password).
		Save(ctx)
	if err != nil {
		log.SetLevel(log.ErrorLevel)
		log.Errorf("something went wrong: %+v", err)

		return nil, err
	}

	messageBody := fmt.Sprintf("Halo %s, Kode Verifikasi pendaftaran di uinvest adalah : 100987", input.Firstname)

	resp, _ := ar.Twillio.SendMessage(input.Phonenumber, messageBody)

	log.Infof("message has been sent with status code %s", resp.Status)

	return newUser, nil
}

// GetDataUser get one data user from database
func (ar *AuthenticationRepository) GetDataUser(ctx context.Context, input model.LoginInput) (*ent.User, error) {
	user, err := ar.DB.User.
		Query().
		Where(user.And(
			user.EmailEQ(input.Email),
			user.PasswordEQ(input.Password),
		)).
		Only(ctx)

	if err != nil {
		log.SetLevel(log.ErrorLevel)
		log.Errorf("something went wrong: %+v", err)

		return nil, err
	}

	messageBody := fmt.Sprintf("Halo %s, Kode Verifikasi untuk login di uinvest adalah : 100987", user.Firstname)

	resp, _ := ar.Twillio.SendMessage(user.Phonenumber, messageBody)
	log.Infof("message has been sent with status code %s", resp.Status)

	return user, nil
}
