package usecases

import (
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
	"go-gc-community/pkg/google"
	"go-gc-community/pkg/token"
	"strings"
)

type User interface {
	//Redirect() (string, error)
	Redirect() (string)
	Account(request *models.UserLoginRequest) (*models.User, string, error)
}

type userUsecase struct {
	ur repositories.User
	t token.Authorization
	g google.Goog
}

func NewUserUsecase(ur repositories.User, t token.Authorization, g google.Goog) *userUsecase {
	return &userUsecase{ur, t, g}
}

func (uu *userUsecase) Redirect() (string) {
	url := uu.g.Redirect()
	return url
}

func (uu *userUsecase) Account(state string, code string) (*models.User, string, int, error) {
	fetched, err := uu.g.Fetch(state, code)
	if err != nil {
		return nil, "", 0, err
	}

	exist, err := uu.ur.Find("email", fetched.Email)
	if err != nil {
		return nil, "", 0, err
	}

	if exist.ID == 0 {
		input := models.User{
			Name: fetched.Name,
			Email: strings.ToLower(fetched.Email),
			RoleId: "1",
			State: "1",
			IsVolunteer: false,
		}

		user, err := uu.ur.Create(&input)
		if err != nil {
			return nil, "", 0, err
		}

		input.AccountNumber = fmt.Sprintf("1%06d", user.ID)
		update, err := uu.ur.Update(&input)
		if err != nil {
			return &input, "", 0, err
		}

		valid, err := uu.ur.First("email", update.Email)
		if err != nil {
			return valid, "", 0, err
		}

		appToken, err := uu.t.Generate(valid.AccountNumber, valid.ID)
		if err != nil {
			return valid, "", 0, err
		}
		
		return valid, appToken, 201, nil
	}

	appToken, err := uu.t.Generate(exist.AccountNumber, exist.ID)
	if err != nil {
		return exist, "", 0, err
	}

	return exist, appToken, 200, nil
}