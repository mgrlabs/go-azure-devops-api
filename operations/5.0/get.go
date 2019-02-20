// Package operations for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/operations/operations/get?view=azure-devops-rest-5.0
package operations

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/mgrlabs/go-azure-devops-api/tools"
)

var apiVersion = "?api-version=5.0"
var baseURI = "https://dev.azure.com/"
var apiPath = "/_apis/operations/"

// OpsStatus returns the current status of a queued operation within Azure DevOps
func OpsStatus(PAT, opsID, azureDevopsOrg string) (status string) {

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Build the API call
	requestURL := baseURI + azureDevopsOrg + apiPath + opsID + apiVersion
	req, err := http.NewRequest("GET", requestURL, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	if err != nil {
		panic(err)
	}

	// Create client call to API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Decode response body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err) // Change this error handling!
	}

	// Extract the operations status from the response
	status = (string(response))
	return status
}
