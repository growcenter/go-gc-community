package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID					int				`json:"id"`
	Name				string			`json:"name"`
	AccountNumber		string			`json:"account_number"`
	CommunityNumber		string			`json:"community_number"`
	Gender				string			`json:"gender"`
	CommunityID			int				`json:"community_id"`
	CommunityName		string			`json:"community_name"`
	Email				string			`json:"email"`
	PhoneNumber			string			`json:"phone_number"`
	Location			string			`json:"location"`
	Address				string			`json:"address"`
	Age					string			`json:"age"`
	IsVolunteer			bool			`json:"is_volunteer"`
	State				string			`json:"state"`
	RoleId				string			`json:"role"`
	CreatedAt			time.Time		`json:"created_at"`
	DeletedAt			sql.NullTime	`json:"deleted_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
}

type (
	UserLoginRequest struct {
		Name			string		`json:"name" validate:"required"`
		Email			string		`json:"email" validate:"required, email"`
	}

	UserLoginResponse struct {
		ResponseCode	string		`json:"responseCode"`
		ResponseMessage	string		`json:"responseMessage"`
		AccountNumber	string		`json:"accountNumber"`
		Token			string		`json:"token"`
		UserID			int			`json:"userID"`
	}
)