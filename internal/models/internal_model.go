package models

type (
	SetRoleRequest struct {
		Email				string					`json:"email" validate:"required, email"`
		AccountNumber 		string					`json:"accountNumber" validate:"required"`
		RoleId				string					`json:"roleId" validate:"required"`
	}
	SetRoleResponse struct {
		ResponseCode		string					`json:"responseCode"`
		ResponseMessage		string					`json:"responseMessage"`
		Name				string					`json:"name"`
		Email				string					`json:"email"`
		AccountNumber		string					`json:"accountNumber"`
		RoleId				string					`json:"roleId"`
	}
)

type (
	RegistrationListResponse struct {
		ResponseCode		string						`json:"responseCode"`
		ResponseMessage		string						`json:"responseMessage"`
		Number				int							`json:"pageNumber"`
		Limit				int							`json:"pageLimit"`
		Sort				string						`json:"pageSort"`
		Data				[]RegistrationListDetail 	`json:"data"`
	}
	RegistrationListDetail struct {
		Name				string						`json:"name"`
		Email				string						`json:"email"`
		AccountNumber		string						`json:"accountNumber"`
		Code				string						`json:"code"`
		EventsId			int							`json:"eventsId"`
		SessionsId			int							`json:"sessionsId"`
		Status				string						`json:"status"`
		BookedBy			string						`json:"bookedBy"`
	
	}
)

type (
	VerifyRegistrationRequest struct {
		Code				string						`json:"code" validate:"required"`
	}
	VerifyRegistrationResponse struct {
		ResponseCode		string						`json:"responseCode"`
		ResponseMessage		string						`json:"responseMessage"`
		Name				string						`json:"name"`
		Email				string						`json:"email"`
		Status				string						`json:"status"`
		Information			EventInformationDetail		`json:"eventInformation"`
		Statistics			EventStatisticsDetail		`json:"eventStatistics"`
	}
	EventInformationDetail struct {
		EventsId			int							`json:"eventsId"`
		EventName			string						`json:"eventName"`
		SessionsId			int							`json:"sessionsId"`
		SessionName			string						`json:"sessionName"`
	}
	EventStatisticsDetail struct {
		AvailableSeats		int							`json:"availableSeats"`
		BookedSeats			int							`json:"bookedSeats"`
		ScannedSeats		int							`json:"scannedSeats"`
		UnscannedSeats		int							`json:"unscannedSeats"`
		TotalRegistration	int							`json:"totalRegistration"`
	}
)