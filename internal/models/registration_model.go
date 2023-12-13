package models

import (
	"database/sql"
	"time"
)

type Registrations struct {
	ID					int				`json:"id"`
	Name				string			`json:"name"`
	Email				string			`json:"email"`
	Code				string			`json:"code"`
	EventsId			int				`json:"events_id"`
	SessionsId			int				`json:"sessions_id"`
	Status				string			`json:"status"`
	BookedBy			string			`json:"booked_by"`
	UpdatedBy			string			`json:"updated_by"`
	AccountNumber		string			`json:"account_number"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
	DeletedAt			sql.NullTime	`json:"deleted_at"`
}

type (
	RegistrationRequest struct {
		EventID				int						`json:"eventId" binding:"required"`
		SessionID			int						`json:"eventSessionId" binding:"required"`
		MainEmail			string					`json:"mainEmail" binding:"required"`
		MainName			string					`json:"mainName" binding:"required"`
		Others				[]OtherBookingRequest	`json:"otherBooking"`
	}
	OtherBookingRequest struct {
		Email				string					`json:"email" binding:"email"`
		Name				string					`json:"name"`
	}
	RegistrationResponse struct {
		ResponseCode		string					`json:"responseCode"`
		ResponseMessage		string					`json:"responseMessage"`
		EventCode			string					`json:"bookedEvent"`
		IsValid				bool					`json:"isValid"`
		SessionID			int						`json:"eventSessionId"`
		EventID				int						`json:"eventId"`
		RequestedSeats		int						`json:"eventRequestedSeats"`
		MainEmail			string					`json:"mainEmail"`
		MainName			string					`json:"mainName"`
		MainAccountNumber	string					`json:"mainAccountNumber"`
		MainBookingCode		string					`json:"mainBookingCode"`
		Others				[]OtherBookingResponse	`json:"otherBooking"`
	}
	OtherBookingResponse struct {
		Email				string					`json:"email"`
		Name				string					`json:"name"`
		BookingCode			string					`json:"bookingCode"`
	}
)

type (
	ViewRegistrationResponse struct {
		ResponseCode		string					`json:"responseCode"`
		ResponseMessage		string					`json:"responseMessage"`
		MainEmail			string					`json:"mainEmail"`
		MainName			string					`json:"mainName"`
		MainStatus			string					`json:"mainStatus"`
		MainAccountNumber	string					`json:"mainAccountNumber"`
		EventName			string					`json:"eventName"`
		SessionName			string					`json:"sessioName"`
		SessionTime			string					`json:"sessionTime"`
		Others				[]OtherRegisResponse	`json:"otherBooking,omitempty"`
	}
	OtherRegisResponse struct {
		Email				string					`json:"email"`
		Name				string					`json:"name"`
		BookingCode			string					`json:"bookingCode"`
		Status				string					`json:"status"`
	}
)

type (
	CancelRegistrationRequest struct {
		Email				string					`json:"email" binding:"required"`
		Code				string					`json:"code"`
	}
)