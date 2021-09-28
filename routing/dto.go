package routing

// Request ...
type Request struct {
	Type       string
	Username   string
	BookingID  string
	Signature  string
	APIVersion string
}

// BillerData ...
type BillerData struct {
	BillerID     int64
	Username     string
	SpecVersion  string
	Password     string
	GenerateMode string
	Status       int64
	Email        string
	Name         string
	Prefix       string
}
