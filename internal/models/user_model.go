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
	Password			string			`json:"password"`
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
		Name			string		`json:"name" binding:"required"`
		Email			string		`json:"email" binding:"required"`
	}

	UserLoginResponse struct {
		ResponseCode	string		`json:"responseCode"`
		ResponseMessage	string		`json:"responseMessage"`
		AccountNumber	string		`json:"accountNumber"`
		Token			string		`json:"token"`
		UserID			int			`json:"userID"`
	}
)

type (
	UserManualRegisterRequest struct {
		Name			string		`json:"name" binding:"required"`
		Identifier		string		`json:"identifier" binding:"required"`
		Password		string		`json:"password" binding:"required"`
	}
	UserManualRegisterResponse struct {
		ResponseCode	string		`json:"responseCode"`
		ResponseMessage	string		`json:"responseMessage"`
		Name			string		`json:"name"`
		Email			string		`json:"email,omitempty"`
		PhoneNumber		string		`json:"phoneNumber,omitempty"`
		Password		string		`json:"password"`
		AccountNumber	string		`json:"accountNumber"`
		Token			string		`json:"token"`
		UserID			int			`json:"userID"`
	}
)

type (
	UserManualLoginRequest struct {
		Identifier		string		`json:"identifier" binding:"required"`
		Password		string		`json:"password" binding:"required"`
	}
	UserManualLoginResponse struct {
		ResponseCode	string		`json:"responseCode"`
		ResponseMessage	string		`json:"responseMessage"`
		Name			string		`json:"name"`
		Email			string		`json:"email,omitempty"`
		PhoneNumber		string		`json:"phoneNumber,omitempty"`
		AccountNumber	string		`json:"accountNumber"`
		Token			string		`json:"token"`
		UserID			int			`json:"userID"`
	}
)

type (
	InquiryUserRequest struct {
		AccountNumber	string			`json:"accountNumber" binding:"required"`
		Additional		AdditionalInfo	`json:"additionalInfo"`
	}
	
	InquiryUserResponse struct {
		ResponseCode	string		`json:"responseCode"`
		ResponseMessage	string		`json:"responseMessage"`
		AccountNumber	string		`json:"accountNumber"`
		Name			string		`json:"name"`
		State			string		`json:"state"`
		Role			string		`json:"role"`
		Email			string		`json:"email"`
	}
)