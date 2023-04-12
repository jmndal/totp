package views

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func GenerateTOTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/totp.html"))
	context := map[string]interface{}{}

	secret := "JBSWY3DPEHPK3PXP" // Example secret from RFC 6238
	totp, err := TOTPGenerator(secret)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("TOTP:", totp)
	}

	tmpl.Execute(w, context)
}

func TOTPGenerator(secret string) (string, error) {
	// Decode the secret from base32
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// Get the current Unix time in seconds
	currentTime := time.Now().Unix()

	// Compute the number of time steps that have elapsed since the Unix epoch
	timeStep := 30 // TOTP time step in seconds
	timeSteps := currentTime / int64(timeStep)

	// Convert the time steps to a byte array
	timeBytes := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		timeBytes[i] = byte(timeSteps & 0xff)
		timeSteps = timeSteps >> 8
	}

	// Compute the HMAC-SHA256 hash of the time bytes using the secret key
	hmacSha256 := hmac.New(sha256.New, secretBytes)
	hmacSha256.Write(timeBytes)
	hash := hmacSha256.Sum(nil)

	// Extract the 4-byte dynamic offset from the last 4 bits of the hash
	offset := hash[len(hash)-1] & 0xf
	code := ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1]) & 0xff) << 16) |
		((int(hash[offset+2]) & 0xff) << 8) |
		(int(hash[offset+3]) & 0xff)

	// Truncate the code to a 6-digit TOTP value
	totp := code % 1000000

	// Convert the TOTP value to a string
	return fmt.Sprintf("%06d", totp), nil
}
