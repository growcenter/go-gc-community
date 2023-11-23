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
	AccountNumber		string			`json:"account_number"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
	DeletedAt			sql.NullTime	`json:"deleted_at"`
}

type (
	RegistrationRequest struct {
		EventID				int						`json:"eventId" validate:"required"`
		SessionID			int						`json:"eventSessionId" validate:"required"`
		MainEmail			string					`json:"mainEmail" validate:"required"`
		MainName			string					`json:"mainName" validate:"required"`
		Others				[]OtherBookingRequest	`json:"otherBooking"`
	}
	OtherBookingRequest struct {
		Email				string					`json:"email"`
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