package views

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"html/template"
	"image/png"
	"net/http"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var secretBase32 = ""

func GenerateTOTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/totp.html"))
	context := map[string]interface{}{}

	data_action := r.FormValue("data_action")
	fmt.Println("data_action: ", data_action)

	if data_action == "GENERATE_KEY" {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Your Organization",
			AccountName: "Username",
			Period:      30,
			SecretSize:  10,
			Algorithm:   otp.AlgorithmSHA256,
		})
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		secretBase32 = key.Secret()
		qrCode, err := key.Image(200, 200)
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		qrCodeBuffer := new(bytes.Buffer)
		err = png.Encode(qrCodeBuffer, qrCode)
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		qrCodeBase64 := base64.StdEncoding.EncodeToString(qrCodeBuffer.Bytes())
		context["generateSecret"] = key.Secret()
		context["qrCode"] = qrCodeBase64
	}

	if data_action == "GENERATE_TOTP" {
		totpCode, err := TOTPGenerator(secretBase32)
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Println("TOTP:", totpCode)
		context["generateTOTP"] = totpCode
	}
	tmpl.Execute(w, context)
}

func ValidateTOTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/totp.html"))
	context := map[string]interface{}{}

	data_action := r.FormValue("data_action")
	fmt.Println("data_action: ", data_action)

	if data_action == "VALIDATE_TOTP" {
		totpCode := r.FormValue("totp_code")
		valid := ValidateTOTPCode(secretBase32, totpCode)
		context["validationResult"] = valid
	}

	tmpl.Execute(w, context)
}

func TOTPGenerator(secret string) (string, error) {
	key, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return key, nil
}

func ValidateTOTPCode(secret string, code string) bool {
	decoded, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return false
	}

	counter := uint64(time.Now().Unix() / 30)
	mac := hmac.New(sha256.New, decoded)
	err = binary.Write(mac, binary.BigEndian, counter)
	if err != nil {
		return false
	}

	expectedCode := hotpTruncate(mac.Sum(nil))
	actualCode, err := base32.StdEncoding.DecodeString(code)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedCode, actualCode)
}

func hotpTruncate(hashedData []byte) []byte {
	offset := int(hashedData[len(hashedData)-1] & 0xf)
	binary := ((int(hashedData[offset]) & 0x7f) << 24) |
		((int(hashedData[offset+1] & 0xff)) << 16) |
		((int(hashedData[offset+2] & 0xff)) << 8) |
		(int(hashedData[offset+3]) & 0xff)
	return []byte(fmt.Sprintf("%06d", binary%1000000))
}
