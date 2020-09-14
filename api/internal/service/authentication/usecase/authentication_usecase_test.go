package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/awesomebusiness/uinvest/internal/model"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationUsecase_RegisterValidation(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockRepository := mock.NewMockRepositoryAuthentication(controller)

	t.Run("register validation should be success", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "+6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=", // encrypted password
			Phonenumber: "+6283834121715",
		}

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(&model.User{}, nil).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, user)
	})

	t.Run("register validation failed insert into database", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "+6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=", // encrypted password
			Phonenumber: "+6283834121715",
		}

		errExpect := fmt.Errorf("something went wrong: failed input to database")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("register validation failed empty input email", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=", // encrypted password
			Phonenumber: "6283834121715",
		}

		errExpect := fmt.Errorf("email should not be empty")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("register validation failed empty input firstname", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "",
			Lastname:    "Pribadi",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=", // encrypted password
			Phonenumber: "6283834121715",
		}

		errExpect := fmt.Errorf("firstname should not be empty")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("register validation failed empty input lastname", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "",
			Password:    "risky123",
			Phonenumber: "6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=", // encrypted password
			Phonenumber: "6283834121715",
		}

		errExpect := fmt.Errorf("lastname should not be empty")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("register validation failed empty input password", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "",
			Phonenumber: "+6283834121715",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "",
			Phonenumber: "+6283834121715",
		}

		errExpect := fmt.Errorf("password should not be empty")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("register validation failed not valid input phonenumber", func(t *testing.T) {
		input := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "risky123",
			Phonenumber: "0838341",
		}

		inputExpect := model.RegisterInput{
			Email:       "riskypribadi@gmail.com",
			Firstname:   "Risky",
			Lastname:    "Pribadi",
			Password:    "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=",
			Phonenumber: "0838341",
		}

		errExpect := fmt.Errorf("phonenumber should not be empty")

		mockRepository.EXPECT().CreateDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.RegisterValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})
}

func TestAuthenticationUsecase_AuthenticationValidation(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockRepository := mock.NewMockRepositoryAuthentication(controller)

	t.Run("authentication validation should be success", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi@gmail.com",
			Password: "risky123",
		}

		inputExpect := model.LoginInput{
			Email:    "riskypribadi@gmail.com",
			Password: "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=",
		}

		mockRepository.EXPECT().GetDataUser(ctx, inputExpect).Return(&model.User{}, nil).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.AuthenticationValidation(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, user)
	})

	t.Run("authentication validation failed user not found", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi21@gmail.com",
			Password: "risky123",
		}

		inputExpect := model.LoginInput{
			Email:    "riskypribadi21@gmail.com",
			Password: "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=",
		}

		errExpect := fmt.Errorf("user not found")

		mockRepository.EXPECT().GetDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.AuthenticationValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("authentication validation failed empty input email", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "",
			Password: "risky123",
		}

		inputExpect := model.LoginInput{
			Email:    "",
			Password: "ywHTbCrZCoa2+o3Z34r7HMAAo5Bj6wU/F9AMvXWEKs8=",
		}

		errExpect := fmt.Errorf("email should not be empty")

		mockRepository.EXPECT().GetDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.AuthenticationValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})

	t.Run("authentication validation failed empty input password", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "riskypribadi21@gmail.com",
			Password: "",
		}

		inputExpect := model.LoginInput{
			Email:    "riskypribadi21@gmail.com",
			Password: "",
		}

		errExpect := fmt.Errorf("password should not be empty")

		mockRepository.EXPECT().GetDataUser(ctx, inputExpect).Return(nil, errExpect).Times(1)

		authenticationUC := NewAuthenticationUsecase(mockRepository)

		user, err := authenticationUC.AuthenticationValidation(ctx, input)

		require.Error(t, err)
		require.Nil(t, user)
	})
}
