package patencode

import (
	b64 "encoding/base64"
)

// PATEncode accepts standard DevOps Personal Access Token and converts
// it to base64 encoded for API calls
func PATEncode(pat string) (PAT string) {

	PAT = b64.StdEncoding.EncodeToString([]byte(":" + pat))
	return PAT
}
