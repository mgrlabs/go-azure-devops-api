package encodepat

// Golang Azure DevOps API
// by mgrlabs
// Encodes a standard Azure DevOps Personal Access Token and returns it encoded for API auth

import (
	b64 "encoding/base64"
)

// PATEncode accepts standard DevOps Personal Access Token and converts
// it to base64 encoded for API calls
func PATEncode(pat string) (encodedPAT string) {

	encodedPAT = b64.StdEncoding.EncodeToString([]byte(":" + pat))
	return encodedPAT
}
