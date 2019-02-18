// Package coreprojectcreate for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/create?view=azure-devops-rest-5.0
package coreprojectcreate

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
	processes "github.com/mgrlabs/go-azure-devops-api/workitemtrackingprocess/processes/5.0"
	gjson "github.com/tidwall/gjson"
)

var apiVersion = "5.0"
var baseURI = "https://dev.azure.com/"
var apiPath = "/_apis/projects?api-version="

// Payload for Project creation - Outer
type Payload struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Capabilities Capabilities `json:"capabilities"`
}

// Capabilities struct - Inner from Payload struct
type Capabilities struct {
	VersionControl  SourceControl     `json:"versioncontrol"`
	ProcessTemplate ProcessTemplateID `json:"processTemplate"`
}

// SourceControl struct - Inner-most from Capabilties struct
type SourceControl struct {
	SourceControlType string `json:"sourceControlType"`
}

// ProcessTemplateID struct - Inner-most from Capabilities struct
type ProcessTemplateID struct {
	TemplateTypeID string `json:"templateTypeId"`
}

// Client struct - Usage TBD
type Client struct {
}

// ProjectResponse struct - Usage TBD
type ProjectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

// CreateProject creates the Azure DevOps project
func CreateProject(PAT, azureDevopsOrg, projectName, workItemProcess, description, versionControl string) ProjectResponse {

	// Call to work item process list function, returns JSON payload containing templates
	processGUID := gjson.Get(processes.ProcessTemplates(PAT, azureDevopsOrg),
		`value.#[name="`+workItemProcess+`"].typeId`)

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Pass Project-specific parms into variable
	payload := Payload{
		Name:        projectName,
		Description: description,
		Capabilities: Capabilities{
			VersionControl: SourceControl{
				SourceControlType: versionControl,
			},
			ProcessTemplate: ProcessTemplateID{
				TemplateTypeID: processGUID.String(),
			},
		},
	}

	// Build JSON Payload
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Build API call
	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payloadJSON))
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	req.Header.Set("Content-Type", "application/json")

	// Request Project creation
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Decode response body
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := ProjectResponse{}
	json.Unmarshal([]byte(response), &data)

	return data
}
