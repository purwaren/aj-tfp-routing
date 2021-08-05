

package main

import (
	"crypto/md5"
	"encoding/hex"

	"log"
	"net/http"
	"routing/request"
	_ "routing/request"
)






func Verification (r *http.Request, billerConfig request.BillerConfig, req request.Request, billerData request.BillerData)(string, error){
	var ack string
	ack = "00"

	//validation type
	if req.Type != "reqpaymentcode" {
		ack = "02"
		return ack, nil
	}
	//validation content type
	if r.Header.Get("Content-Type") != "text/xml" || r.Header.Get("Content-Type") != "application/xml" {
		ack = "03"
		return ack, nil

	}

	//validation biller status
	if billerData.Status == 1 {
		ack ="00"
	} else {
		ack = "02"
		return ack, nil
	}

	//verification signatur

	hasher := md5.New()
	hasher.Write([]byte(billerConfig.Username+ billerConfig.Secret + req.BookingID))

	if req.Signature != hex.EncodeToString(hasher.Sum(nil)) {
		ack = "01"
		return ack, nil
	}

	return ack, nil
}
func handleRequests() {
	
	log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
	handleRequests()
}
