// Package coreproject for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/create?view=azure-devops-rest-5.0
package coreproject

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	operations "github.com/mgrlabs/go-azure-devops-api/operations/5.0"
	tools "github.com/mgrlabs/go-azure-devops-api/tools"
	processes "github.com/mgrlabs/go-azure-devops-api/workitemtrackingprocess/processes/5.0"
	gjson "github.com/tidwall/gjson"
)

// Payload for Project creation - Outer
type Payload struct {
	Name         string       `json:"name"`
	Visibility   string       `json:"visibility"`
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
func CreateProject(PAT, azureDevopsOrg, projectName, workItemProcess, description, versionControl, visibility string) ProjectResponse {

	var apiVersion = "5.0"
	var baseURI = "https://dev.azure.com/"
	var apiPath = "/_apis/projects?api-version="

	// Call to work item process list function, returns JSON payload containing templates
	processGUID := gjson.Get(processes.ProcessTemplates(PAT, azureDevopsOrg),
		`value.#[name="`+workItemProcess+`"].typeId`)

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(PAT)

	// Pass Project-specific parms into variable
	payload := Payload{
		Name:        projectName,
		Visibility:  visibility,
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
	data := ProjectResponse{}
	if err != nil {
		data.ID = "1"
		data.Status = "ERROR: Could not build JSON payload!"
		return data
	}

	// Build API call
	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		data.ID = "1"
		data.Status = "ERROR: Could not build API Call!"
		return data
	}
	req.Header.Set("Authorization", "Basic "+encodedPAT)
	req.Header.Set("Content-Type", "application/json")

	// Request Project creation
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data.ID = "1"
		data.Status = "ERROR: Could not connect to Azure DevOps API!"
		return data
	}

	// Decode response body
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
				time.Sleep(2 * time.Second)
			case "succeeded":
				time.Sleep(1 * time.Second)
				g := gjson.Get(ListProjects(PAT, azureDevopsOrg), `value.#[name="`+projectName+`"].id`)
				data.ID = g.String()
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
