package usecases

import (
	"errors"
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
	"go-gc-community/pkg/authorization"
	"go-gc-community/pkg/google"
	"go-gc-community/pkg/hash"
	"go-gc-community/pkg/validate"
	"strings"
)

type User interface {
	Redirect() (string)
	Account(state string, code string) (*models.User, string, int, error)
	Inquire(request *models.InquiryUserRequest) (*models.User, error)
	ManualRegister(request *models.UserManualRegisterRequest) (*models.User, error)
	ManualLogin(request *models.UserManualLoginRequest) (*models.User, error)
}

type userUsecase struct {
	ur repositories.User
	a authorization.Authorization
	g google.Goog
	s []byte
}

func NewUserUsecase(ur repositories.User, a authorization.Authorization, g google.Goog, s []byte) *userUsecase {
	return &userUsecase{ur, a, g, s}
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
			RoleId: "01",
			State: "1",
			IsVolunteer: false,
		}

		user, err := uu.ur.Create(&input)
		if err != nil {
			return nil, "", 0, err
		}

		input.AccountNumber = fmt.Sprintf("1%09d", user.ID)
		update, err := uu.ur.Update(&input)
		if err != nil {
			return &input, "", 0, err
		}

		valid, err := uu.ur.First("email", update.Email)
		if err != nil {
			return valid, "", 0, err
		}

		appToken, err := uu.a.Generate(valid.AccountNumber, valid.ID)
		if err != nil {
			return valid, "", 0, err
		}
		
		return valid, appToken, 201, nil
	}

	appToken, err := uu.a.Generate(exist.AccountNumber, exist.ID)
	if err != nil {
		return exist, "", 0, err
	}

	return exist, appToken, 200, nil
}

func (uu *userUsecase) Inquire(request *models.InquiryUserRequest) (*models.User, error) {
	user, err := uu.ur.First("account_number", request.AccountNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) ManualRegister(request *models.UserManualRegisterRequest) (*models.User, string, error) {
	isEmail := validate.Email(strings.ToLower(request.Identifier))
	isPhone := validate.PhoneNumber(request.Identifier)
	
	isExist, err := uu.ur.FindMultipleExact("phone_number", "email", strings.ToLower(request.Identifier))
	if err != nil {
		return nil, "", err
	}

	if isExist.ID != 0 || strings.ToLower(request.Identifier) == strings.ToLower(isExist.Email) || strings.ToLower(request.Identifier) == strings.ToLower(isExist.PhoneNumber) {
		return nil, "", errors.New("account already registered")
	}
	
	if !isEmail {
		if !isPhone {
			return nil, "", errors.New("should input either valid email or phone number")	
		}
		
		salted := append([]byte(request.Password), uu.s...)
		password, err := hash.Generate(salted)
		if err != nil {
			return nil, "", err
		}

		input := models.User{
			Name: request.Name,
			PhoneNumber: strings.ToLower(request.Identifier),
			RoleId: "01",
			State: "1",
			IsVolunteer: false,
			Password: password,
		}
	
		user, err := uu.ur.Create(&input)
		if err != nil {
			return nil, "", err
		}
	
		input.AccountNumber = fmt.Sprintf("1%09d", user.ID)
		update, err := uu.ur.Update(&input)
		if err != nil {
			return nil, "", err
		}
	
		/*valid, err := uu.ur.First("email", update.Email)
		if err != nil {
			return nil, "", err
		}*/

		valid, err := uu.ur.First("id", update.ID)
		if err != nil {
			return nil, "", err
		}

		appToken, err := uu.a.Generate(valid.AccountNumber, valid.ID)
		if err != nil {
			return nil, "", err
		}

		return valid, appToken, nil
	}
	
	salted := append([]byte(request.Password), uu.s...)
	password, err := hash.Generate(salted)
	if err != nil {
		return nil, "", err
	}

	input := models.User{
		Name: request.Name,
		Email: strings.ToLower(request.Identifier),
		RoleId: "01",
		State: "1",
		IsVolunteer: false,
		Password: password,
	}

	user, err := uu.ur.Create(&input)
	if err != nil {
		return nil, "", err
	}

	input.AccountNumber = fmt.Sprintf("1%09d", user.ID)
	update, err := uu.ur.Update(&input)
	if err != nil {
		return nil, "", err
	}

	/*valid, err := uu.ur.First("email", update.Email)
	if err != nil {
		return nil, "", err
	}*/

	valid, err := uu.ur.First("id", update.ID)
	if err != nil {
		return nil, "", err
	}

	appToken, err := uu.a.Generate(valid.AccountNumber, valid.ID)
	if err != nil {
		return nil, "", err
	}

	return valid, appToken, nil
}

func (uu *userUsecase) ManualLogin(request *models.UserManualLoginRequest) (*models.User, string, error) {
	user, err := uu.ur.FindMultipleExact("phone_number", "email", strings.ToLower(request.Identifier))
	if err != nil {
		return nil, "", err
	}
	
	if user.ID == 0 {
		return nil, "", errors.New("user has not registered yet")
	}
	
	salted := append([]byte(request.Password), uu.s...)
	err = hash.Validate(user.Password, string(salted))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	appToken, err := uu.a.Generate(user.AccountNumber, user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, appToken, nil
}