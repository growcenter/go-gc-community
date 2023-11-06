package models

type User struct {
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	AccountNumber	string		`json:"account_number"`
	CommunityNumber	string		`json:"community_number"`
	Gender			string		`json:"gender"`
	CommunityID		int			`json:"community_id"`
	CommunityName	string		`json:"community_name"`
	Email			string		`json:"email"`
	PhoneNumber		string		`json:"phone_number"`
	Password		string		`json:"password"`
	Location		string		`json:"location"`
	Category		string		`json:"category"`
	GeneralEvent	string		`json:"general_event"`
	IsVolunteer		bool		`json:"is_volunteer"`
	VolunteerEvent	string		`json:"volunteer_event"`
	State			string		`json:"state"`
}