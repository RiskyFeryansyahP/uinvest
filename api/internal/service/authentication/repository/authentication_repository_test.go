package repository

import (
	"context"
	"testing"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/require"
)

func TestAuthenticationRepository(t *testing.T) {
	ctx := context.Background()

	client, err := sqliteConnection()

	require.NoError(t, err)

	err = client.Schema.Create(ctx)

	require.NoError(t, err)

	authenticationRepo := NewAuthenticationRepository(client)

	// Testing CreateDataUser
	t.Run("create data user should success", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "6283834121715",
		}

		newUser, err := authenticationRepo.CreateDataUser(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, newUser)
	})

	t.Run("create data failed there is an empty field", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "",
		}

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

		user, err := authenticationRepo.GetDataUser(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("get data failed there is an empty input", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "",
			Password: "",
		}

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
