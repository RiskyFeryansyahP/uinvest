package resolver

import (
	"context"
	"fmt"
	"testing"

	"github.com/awesomebusiness/uinvest/ent"
	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	mockAuthenticationUC := mock.NewMockUsecaseAuthentication(ctrl)

	t.Run("login should be success", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi@gmail.com",
			Password: "risky123",
		}

		mockAuthenticationUC.EXPECT().AuthenticationValidation(ctx, input).Times(1).Return(&ent.User{}, nil)

		resolverAuthentication := &Resolver{
			AuthenticationUC: mockAuthenticationUC,
		}

		user, err := resolverAuthentication.Query().Login(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, user)
	})

	t.Run("login failed error in validation empty email", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "",
			Password: "risky123",
		}

		errExpect := fmt.Errorf("email is not valid")

		mockAuthenticationUC.EXPECT().AuthenticationValidation(ctx, input).Times(1).Return(nil, errExpect)

		resolverAuthentication := &Resolver{
			AuthenticationUC: mockAuthenticationUC,
		}

		user, err := resolverAuthentication.Query().Login(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	mockAuthenticationUC := mock.NewMockUsecaseAuthentication(ctrl)

	t.Run("register should be success", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "62838341217119",
		}

		mockAuthenticationUC.EXPECT().RegisterValidation(ctx, input).Times(1).Return(&ent.User{}, nil)

		resolverAuthentication := &Resolver{
			AuthenticationUC: mockAuthenticationUC,
		}

		newUser, err := resolverAuthentication.Query().Register(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, newUser)
	})

	t.Run("register failed empty input firstname", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "62838341217119",
		}

		errExpect := fmt.Errorf("firstname cant be empty")

		mockAuthenticationUC.EXPECT().RegisterValidation(ctx, input).Times(1).Return(nil, errExpect)

		resolverAuthentication := &Resolver{
			AuthenticationUC: mockAuthenticationUC,
		}

		newUser, err := resolverAuthentication.Query().Register(ctx, input)

		require.Error(t, err)
		require.Nil(t, newUser)
	})
}
