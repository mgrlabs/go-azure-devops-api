// Package tools for Azure DevOps Go SDK
// Encodes a standard Azure DevOps Personal Access Token and returns it encoded for API auth
// by mgrlabs - github.com/mgrlabs
//
package tools

import (
	b64 "encoding/base64"
)

// PATEncode accepts Personal Access Token and converts it to base64 encoded for API calls
func PATEncode(pat string) (encodedPAT string) {
	encodedPAT = b64.StdEncoding.EncodeToString([]byte(":" + pat))
	return encodedPAT
}
