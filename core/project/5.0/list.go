package coreproject

import (
	"bytes"
	"io/ioutil"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

// ProjectList will return a list of Azure DevOps projects for a given organization
func ProjectList(PAT, azureDevopsOrg string) (responseJSON string) {

	// API-specific settings
	var apiVersion = "5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects?api-version="

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Build the API Call
	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("GET", requestURL, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	if err != nil {
		panic(err)
	}

	// Make the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Decode the response body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Convert body to raw JSON response
	responseJSON = string(response)
	return responseJSON
}
