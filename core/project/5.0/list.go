package coreproject

import (
	"bytes"
	"io/ioutil"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

// ProjectList will return a list of Azure DevOps projects for a given organization
func ProjectList(PAT, azureDevopsOrg string) (projectPayload string) {

	var apiVersion = "5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects?api-version="

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("GET", requestURL, bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	projectPayload = string(response)
	return projectPayload
}
