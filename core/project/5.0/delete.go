// Package coreproject for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/delete?view=azure-devops-rest-5.0
package coreproject

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	operations "github.com/mgrlabs/go-azure-devops-api/operations/5.0"
	tools "github.com/mgrlabs/go-azure-devops-api/tools"
	gjson "github.com/tidwall/gjson"
)

// DeleteProject creates the Azure DevOps project
func DeleteProject(PAT, azureDevopsOrg, projectGUID string) ProjectResponse {

	// API-specific settings
	var apiVersion = "?api-version=5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects/"

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)
	data := ProjectResponse{}

	// Build the API Call
	requestURL := baseURI + azureDevopsOrg + apiPath + projectGUID + apiVersion
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		data.ID = "1"
		data.Status = "ERROR: Could not read API response!"
		return data
	}

	json.Unmarshal([]byte(body), &data)

	if gjson.Get(string(body), "message").Exists() {
		r := gjson.Get(string(body), "message")
		data.ID = "1"
		data.Status = r.String()
		return data
	} else {
		var s string
		for s != "succeeded" {
			r := operations.OpsStatus(PAT, data.ID, azureDevopsOrg)
			s = gjson.Get(r, "status").String()
			switch s {
			case "inProgress", "queued", "notSet":
				time.Sleep(1 * time.Second)
			case "succeeded":
				data.ID = ""
				return data
			case "failed", "cancelled":
				data.ID = "1"
				data.Status = "Unknown Error has occured!"
				return data
			}
		}
	}
	data.ID = "1"
	data.Status = "Unknown Error has occured!"
	return data
}
