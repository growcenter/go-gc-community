package models

import (
	"database/sql"
	"time"
)

type Events struct {
	ID					int				`json:"id"`
	Name				string			`json:"name"`
	Code				string			`json:"code"`
	Status				string			`json:"status"`
	Description			string			`json:"description"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
	DeletedAt			sql.NullTime	`json:"deleted_at"`
}

type (
	GetEventResponse struct {
		ResponseCode		string					`json:"responseCode"`
		ResponseMessage		string					`json:"responseMessage"`
		EventCount			int						`json:"eventCount"`
		Events				[]EventResponseDetail	`json:"events"`
	}
	EventResponseDetail struct {
		EventID				int			`json:"eventId"`
		EventName			string		`json:"eventName"`
		EventDescription	string		`json:"eventDescription"`
		EventCode			string		`json:"eventCode"`
		Status				string		`json:"eventStatus"`
	}
)
