package coreprojectcreate

// https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/create?view=azure-devops-rest-5.0
//

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tools "github.com/mgrlabs/go-azure-devops-api/tools"
)

var apiVersion = "5.0"
var baseURL = "https://dev.azure.com/"
var apiURL = "/_apis/projects?api-version=" + apiVersion

//Client for manage azure devops organization
type Client struct {
}

type ProjectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

// CreateProject creates the Azure DevOps project
func CreateProject(pat, organization, projectName, workItemProcess, description, versionControl string) ProjectResponse {

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(pat)

	processTemplates := map[string]string{
		"Agile": "adcc42ab-9882-485e-a3ed-7678f01f66bc",
		"Scrum": "6b724908-ef14-45cf-84f8-768b5384da45",
		"CMMI":  "27450541-8e31-4150-9947-dc59f998fc01",
	}

	jsonStr := []byte("{ \"name\": \"" + projectName + "\", \"description\": \"" + description + "\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"" + versionControl + "\"}, \"processTemplate\": {  \"templateTypeId\": \"" + processTemplates[workItemProcess] + "\" }}}")
	callURL := baseURL + organization + apiURL
	req, err := http.NewRequest("POST", callURL, bytes.NewBuffer(jsonStr))
	basic := "Basic " + encodedPAT

	req.Header.Set("Authorization", basic)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	data := ProjectResponse{}
	json.Unmarshal([]byte(responseData), &data)

	return data

}
