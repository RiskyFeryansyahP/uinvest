package repository

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/mock"
	"github.com/golang/mock/gomock"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/require"
)

func TestAuthenticationRepository(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	client, err := sqliteConnection()

	require.NoError(t, err)

	err = client.Schema.Create(ctx)

	require.NoError(t, err)

	// Testing CreateDataUser
	t.Run("create data user should success", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "+6283834121715",
		}

		twillioMock := mock.NewMockTwillioMessage(ctrl)
		twillioMock.EXPECT().SendMessage(input.Phonenumber, "Halo Risky, Kode Verifikasi pendaftaran di uinvest adalah : 100987").Times(1).Return(&http.Response{}, nil)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)

		newUser, err := authenticationRepo.CreateDataUser(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, newUser)
	})

	t.Run("failed send msg when success create data", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "+6283834121715",
		}

		errTwillio := fmt.Errorf("failed request http into twillio")

		twillioMock := mock.NewMockTwillioMessage(ctrl)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)

		twillioMock.EXPECT().SendMessage(input.Phonenumber, "Halo Risky, Kode Verifikasi pendaftaran di uinvest adalah : 100987").Times(1).Return(nil, errTwillio)

		newUser, err := authenticationRepo.CreateDataUser(ctx, input)

		require.Error(t, err)
		require.Nil(t, newUser)
	})

	t.Run("create data failed there is an empty field", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "",
		}

		errTwillio := fmt.Errorf("send sms failed, invalid phone number")

		twillioMock := mock.NewMockTwillioMessage(ctrl)
		twillioMock.EXPECT().SendMessage(nil, nil).Times(1).Return(nil, errTwillio)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)

		newUser, err := authenticationRepo.CreateDataUser(ctx, input)

		require.Error(t, err)
		require.Nil(t, newUser)
	})
	// End testing CreateDataUser

	// Testing GetDataUser
	t.Run("get data user should be success", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi@gmail.com",
			Password: "risky123",
		}

		twillioMock := mock.NewMockTwillioMessage(ctrl)
		twillioMock.EXPECT().SendMessage("+6283834121715", "Halo Risky, Kode Verifikasi untuk login di uinvest adalah : 100987").Times(1).Return(&http.Response{}, nil)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)

		user, err := authenticationRepo.GetDataUser(ctx, input)

		require.NoError(t, err)
		require.Equal(t, "Risky", user.Firstname)
		require.Equal(t, "Pribadi", user.Lastname)
	})

	t.Run("get data failed email or password not correct", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi2@gmail.com",
			Password: "risky12",
		}

		errTwillio := fmt.Errorf("failed invalid phone number")

		twillioMock := mock.NewMockTwillioMessage(ctrl)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)
		twillioMock.EXPECT().SendMessage(nil, nil).Times(1).Return(nil, errTwillio)

		user, err := authenticationRepo.GetDataUser(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("get data failed there is an empty input", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "",
			Password: "",
		}

		errTwillio := fmt.Errorf("failed invalid phone number")

		twillioMock := mock.NewMockTwillioMessage(ctrl)

		authenticationRepo := NewAuthenticationRepository(client, twillioMock)
		twillioMock.EXPECT().SendMessage(nil, nil).Times(1).Return(nil, errTwillio)

		user, err := authenticationRepo.GetDataUser(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})
}

func sqliteConnection() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	return client, nil
}
