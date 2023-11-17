package google

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-gc-community/internal/models"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Google interface {
	Redirect() (string)
	//Redirect() (string, error)
	Fetch(state, code string) (*models.GoogleUser, error)
}

type Goog struct {
	authState string
	clientId string
	clientSecret string
	redirectUrl string
}

var oauth = &oauth2.Config{}

func NewGoogle(authState, clientId, clientSecret, redirectUrl string) (*Goog, error) {
	if authState == "" {
		return nil, errors.New("empty signing key")
	}

	return &Goog{authState: authState, clientId: clientId, clientSecret: clientSecret, redirectUrl: redirectUrl}, nil
}

func (g *Goog) Redirect() (string) {
	oauth = &oauth2.Config{
		ClientID:     g.clientId,
		ClientSecret: g.clientSecret,
		RedirectURL:  g.redirectUrl,
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
	
	url := oauth.AuthCodeURL(g.authState)
	return url
}

func (g *Goog) Fetch(state, code string) (*models.GoogleUser, error) {
	if state != g.authState {
		return nil, errors.New("state are not the same")
	}

	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := oauth.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	byteData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	var data map[string]interface{}
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	user := &models.GoogleUser{
		Email: data["email"].(string),
		Name: data["name"].(string),
	}

	fmt.Println(fmt.Sprintf("email: %s\nname: %s", user.Email, user.Name))

	return user, nil
}