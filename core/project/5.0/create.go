package projectcreate

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

//Client for manage azure devops organization
type Client struct {
}

type ProjectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

// CreateProject creates the Azure DevOps project
func CreateProject(pat, organization, projectName, workItemProcess, description string) ProjectResponse {

	// Call to PAT encode function
	encodedPAT := tools.PATEncode(pat)

	var jsonFormat = "{ \"name\": \"" + projectName + "\", \"description\": \"" + description + "\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"Git\"}, \"processTemplate\": {  \"templateTypeId\": \"" + workItemProcess + "\" }}}"

	var jsonStr = []byte(jsonFormat)

	var baseURL = "https://dev.azure.com/" + organization + "/_apis/projects?api-version=5.0"

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonStr))

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
