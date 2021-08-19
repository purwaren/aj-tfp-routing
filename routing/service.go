package routing

import (
	"crypto/md5"
	"encoding/hex"
)

// Verification ...
func Verification(req Request, billerConfig BillerConfig, billerData BillerData) (string, error) {
	ack := "00"

	// Validate content type
	if req.ContentType != "text/xml" && req.ContentType != "application/xml" {
		ack = "03"
		return ack, nil

	}

	// Validate type
	if req.Type != "reqpaymentcode" && req.Type != "reqtrxstatus" {
		ack = "02"
		return ack, nil
	}

	// Validate biller status
	if billerData.Status != 1 {
		ack = "XX"
		return ack, nil
	}

	// Validate api version
	if req.APIVersion != billerConfig.SpecFormat {
		ack = "XX"
		return ack, nil
	}

	// Validate signature
	hasher := md5.New()
	hasher.Write([]byte(req.Username + billerConfig.Secret + req.BookingID))

	if req.Signature != hex.EncodeToString(hasher.Sum(nil)) {
		ack = "01"
		return ack, nil
	}

	return ack, nil
}
