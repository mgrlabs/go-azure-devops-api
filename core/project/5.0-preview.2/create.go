// Package coreprojectcreate for Azure DevOps Go SDK
// by mgrlabs - github.com/mgrlabs
//
// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/create?view=azure-devops-rest-5.0
package coreprojectcreate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

var apiVersion = "5.0-preview.2"
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

// Payload for Project creation - END

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
func CreateProject(pat, azureDevopsOrg, projectName, workItemProcess, description, versionControl string) ProjectResponse {

	// Manual mapping for work process templates - To be replaced by API call
	processTemplates := map[string]string{
		"Agile": "adcc42ab-9882-485e-a3ed-7678f01f66bc",
		"Scrum": "6b724908-ef14-45cf-84f8-768b5384da45",
		"Basic": "b8a3a935-7e91-48b8-a94c-606d37c3e9f2",
		"CMMI":  "27450541-8e31-4150-9947-dc59f998fc01",
	}

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(pat)

	// Pass Project-specific parms into variable
	payload := Payload{
		Name:        projectName,
		Description: description,
		Capabilities: Capabilities{
			VersionControl: SourceControl{
				SourceControlType: versionControl,
			},
			ProcessTemplate: ProcessTemplateID{
				TemplateTypeID: processTemplates[workItemProcess],
			},
		},
	}

	// Build JSON Payload
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Hello")
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
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := ProjectResponse{}
	json.Unmarshal([]byte(responseData), &data)

	return data
}
