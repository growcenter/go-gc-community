package models

type AdditionalInfo struct {
	Channel		string		`json:"channel"`
	Endpoint	string		`json:"source"`
}

type Pagination struct {
	PageNumber	int
	PageLimit	int
	Offset		int
	SortOrder	string
}