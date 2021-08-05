package request


type Request struct {
	Type 	string	`json:"type"`
	Username	string `json:"username"`
	BookingID	string	`json:"bookingid"`
	Signature 	string	`json:"signature"`
}

type BillerConfig struct {
	BillerID	string 	`json:"billerid"`
	Username	string `json:"username"`
	SpecFormat	string	`json:"specformat"`
	Secret		string	`json:"secret"`
	GenerateMode	string	`json:"generatemode"`
}

type BillerData struct {
	BillerID	string 	`json:"billerid"`
	Status		int64	`json:"status"`
	Email		string	`json:"email"`
	Name		string	`json:"name"`
	Prefix		string	`json:"prefix"`
}