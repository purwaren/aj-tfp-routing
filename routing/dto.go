package routing

// Request ...
type Request struct {
	Type      string
	Username  string
	BookingID string
	Signature string
}

// BillerConfig ...
type BillerConfig struct {
	BillerID     string
	Username     string
	SpecFormat   string
	Secret       string
	GenerateMode string
}

// BillerData ...
type BillerData struct {
	BillerID string
	Status   int64
	Email    string
	Name     string
	Prefix   string
}
