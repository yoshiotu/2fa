// 2fa provides a simple end-to-end example of server side support for google authenticator.
// It sets up a secret for a single user, generates a QR code as a PNG file that the user
// can scan into Google Authenticator, and then prompts the user for a token that the user
// copies from the Authenticator app. We validate the token and print out whether it is valid or not.
package main

import (
	// "crypto/rand"
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"net/url"

	qr "code.google.com/p/rsc/qr"
	dgoogauth "github.com/dgryski/dgoogauth"
)

const (
	qrFilename = "/tmp/qr.png"
)

func main() {
	// Example secret from here:
	// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
	secret := []byte{'H', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF}

	// Generate random secret instead of using the test value above.
	// secret := make([]byte, 6)
	// _, err := rand.Read(secret)
	// if err != nil {
	//	panic(err)
	// }

	secretBase32 := base32.StdEncoding.EncodeToString(secret)

	account := "user@example.com"

	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		panic(err)
	}

	URL.Path += "/NameOfMyService:" + account
	params := url.Values{}
	params.Add("secret", secretBase32)
	params.Add("issuer", "NameOfMyService")

	URL.RawQuery = params.Encode()

	fmt.Printf("URL is %s\n", URL.String())

	code, err := qr.Encode(URL.String(), qr.Q)
	if err != nil {
		panic(err)
	}
	b := code.PNG()
	err = ioutil.WriteFile(qrFilename, b, 0600)
	if err != nil {
		panic(err)
	}

	fmt.Printf("QR code is in %s. Please scan it into Google Authenticator app.\n", qrFilename)

	var token string
	fmt.Printf("Please enter the token value: ")
	fmt.Scanln(&token)

	otpc := &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  3,
		HotpCounter: 0,
	}

	val, err := otpc.Authenticate(token)
	if err != nil {
		panic(err)
	}

	if val {
		fmt.Printf("Got true val\n")
	} else {
		fmt.Printf("Got false val\n")
	}
	return
}
