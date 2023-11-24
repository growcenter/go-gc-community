package models

import (
	"database/sql"
	"time"
)

type Sessions struct {
	ID					int				`json:"id"`
	EventsId			int				`json:"events_id"`
	Name				string			`json:"name"`
	Status				string			`json:"status"`
	Description			string			`json:"description"`
	Time				string			`json:"time"`
	MaxSeating			int				`json:"max_seating"`
	AvailableSeats		int				`json:"available_seats"`
	BookedSeats			int				`json:"booked_seats"`
	ScannedSeats		int				`json:"scanned_seats"`
	UnscannedSeats		int				`json:"unscanned_seats"`
	TotalRegistration	int				`json:"total_registration"`
	OpenRegistration	time.Time		`json:"open_registration"`
	ClosedRegistration	time.Time		`json:"closed_registration"`	
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
	DeletedAt			sql.NullTime	`json:"deleted_at"`
}

type (
	GetSessionResponse struct {
		ResponseCode		string					`json:"responseCode"`
		ResponseMessage		string					`json:"responseMessage"`
		EventName			string					`json:"eventName"`
		EventId				int						`json:"eventId"`
		SessionCount		int						`json:"sessionCount"`
		CurrentTime			time.Time				`json:"currentTime"`
		Events				[]SessionResponseDetail	`json:"events"`
	}
	SessionResponseDetail struct {
		Name				string		`json:"eventSessionName"`
		SessionID			int			`json:"eventSessionId"`
		EventID				int			`json:"eventId"`
		Description			string		`json:"eventSessionDescription"`
		Time				string		`json:"eventSessionTime"`
		MaxSeating			int			`json:"eventSessionMaxSeating"`
		Status				string		`json:"eventSessionStatus"`
		IsUserValid			bool		`json:"isUserValid"`
		OpenRegistration	time.Time	`json:"openRegistration"`
		ClosedRegistration	time.Time	`json:"closedRegistration"`
	}
)