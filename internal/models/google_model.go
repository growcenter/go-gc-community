package models

type GoogleUser struct {
	Id			string
	Email		string
	Verified	bool
	Name		string
	GivenName	string
	FamilyName	string
	Picture		string
	Locale		string
}