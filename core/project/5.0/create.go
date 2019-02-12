package azuredevopsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Client for manage azure devops organization
type Client struct {
}

type ProjectResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

func CreateProject(pat string, organization string, projectName string) ProjectResponse {

	var jsonFormat = "{ \"name\": \"" + projectName + "\", \"description\": \"Frabrikam travel app for Windows Phone\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"Git\"}, \"processTemplate\": {  \"templateTypeId\": \"6b724908-ef14-45cf-84f8-768b5384da45\" }}}"

	var jsonStr = []byte(jsonFormat)

	var baseUrl = "https://dev.azure.com/" + organization + "/_apis/projects?api-version=5.0"

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonStr))

	basic := "Basic " + pat

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
