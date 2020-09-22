package repository

import (
	"context"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/ent/user"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication"
	"github.com/awesomebusiness/uinvest/pkg/twillio"

	log "github.com/sirupsen/logrus"
)

// AuthenticationRepository is repository that handle user authentication
type AuthenticationRepository struct {
	DB      *ent.Client
	Twillio twillio.TwillioMessage
}

// NewAuthenticationRepository create new repository with connected to database
func NewAuthenticationRepository(db *ent.Client, twillioClient twillio.TwillioMessage) authentication.RepositoryAuthentication {
	return &AuthenticationRepository{
		DB:      db,
		Twillio: twillioClient,
	}
}

// CreateDataUser create new user profile to database
func (ar *AuthenticationRepository) CreateDataUser(ctx context.Context, input model.RegisterInput) (*model.User, error) {
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

	resp, otp, _ := ar.Twillio.SendOTP(input.Phonenumber, input.Firstname)

	log.Infof("message has been sent with status code %s", resp.Status)

	resultNewUser := &model.User{
		ID:          newUser.ID,
		Email:       newUser.Email,
		Firstname:   newUser.Firstname,
		Lastname:    newUser.Lastname,
		Password:    newUser.Password,
		Phonenumber: newUser.Phonenumber,
		Otp:         otp,
	}

	return resultNewUser, nil
}

// GetDataUser get one data user from database
func (ar *AuthenticationRepository) GetDataUser(ctx context.Context, input model.LoginInput) (*model.User, error) {
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

	resp, otp, _ := ar.Twillio.SendOTP(user.Phonenumber, user.Firstname)

	log.Infof("message has been sent with status code %s", resp.Status)

	resultUser := &model.User{
		ID:          user.ID,
		Email:       user.Email,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Password:    user.Password,
		Phonenumber: user.Phonenumber,
		Otp:         otp,
	}

	return resultUser, nil
}
