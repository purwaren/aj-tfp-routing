package routing

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

// Verification ...
func Verification(header *http.Header, req Request, billerConfig BillerConfig, billerData BillerData) (string, error) {
	ack := "00"

	//validation type
	if req.Type != "reqpaymentcode" {
		ack = "02"
		return ack, nil
	}

	//validation content type
	if header.Get("Content-Type") != "text/xml" || header.Get("Content-Type") != "application/xml" {
		ack = "03"
		return ack, nil

	}

	//validation biller status
	if billerData.Status != 1 {
		ack = "02"
		return ack, nil
	}

	//verification signature
	hasher := md5.New()
	hasher.Write([]byte(billerConfig.Username + billerConfig.Secret + req.BookingID))

	if req.Signature != hex.EncodeToString(hasher.Sum(nil)) {
		ack = "01"
		return ack, nil
	}

	return ack, nil
}
