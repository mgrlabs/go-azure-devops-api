// Package coreproject for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.0
package coreproject

import (
	"bytes"
	"io/ioutil"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

// GetProject returns the properties of a specific project
func GetProject(PAT, azureDevopsOrg, projectGUID string) (data string) {

	// API-specific settings
	var apiVersion = "?api-version=5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects/"

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Build the API Call
	requestURL := baseURI + azureDevopsOrg + apiPath + projectGUID + apiVersion
	req, err := http.NewRequest("GET", requestURL, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	if err != nil {
		panic(err) // Need to update the error handling for this
	}

	// Make the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err) // Need to update the error handling for this
	}
	// Decode the response body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err) // Need to update the error handling for this
	}

	// Convert body to raw JSON response
	data = "test"
	println(response)
	return data
}
