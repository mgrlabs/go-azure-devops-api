// Package coreproject for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/create?view=azure-devops-rest-5.0
package coreproject

import (
	"bytes"
	"io/ioutil"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
	gjson "github.com/tidwall/gjson"
)

// DeleteProject creates the Azure DevOps project
func DeleteProject(PAT, azureDevopsOrg, projectName string) (responseJSON string) {

	var apiVersion = "?api-version=5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects/"

	// Extract the GUID of the project
	projectGUID := gjson.Get(ListProjects(PAT, azureDevopsOrg), `value.#[name="`+projectName+`"].id`)

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Build the API Call
	requestURL := baseURI + azureDevopsOrg + apiPath + projectGUID.String() + apiVersion
	req, err := http.NewRequest("DELETE", requestURL, bytes.NewBuffer(nil))
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
	responseJSON = string(response)
	return responseJSON
}
