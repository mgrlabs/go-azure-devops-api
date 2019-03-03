// Package processes for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/processes/processes/list?view=azure-devops-rest-5.0
package processes

import (
	"io/ioutil"
	"log"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

// ProcessTemplates function returns the work item process templates as a JSON payload
func ProcessTemplates(PAT, azureDevopsOrg string) (responseJSON string) {

	// API-specific settings
	var apiVersion = "5.0-preview.2"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/work/processes?api-version="

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Build API call
	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("Authorization", "Basic "+encodedPAT)

	// Queue project creation
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Close the response body
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	responseJSON = string(response)
	return responseJSON
}
