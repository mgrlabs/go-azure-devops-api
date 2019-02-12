package devopsprojectcreate

import (
	"bytes"
	b64 "encoding/base64"
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

// Comment
func CreateProject(pat string, organization string, projectName string) ProjectResponse {

	patEncoded := b64.StdEncoding.EncodeToString([]byte(":" + pat))

	var jsonFormat = "{ \"name\": \"" + projectName + "\", \"description\": \"Frabrikam travel app for Windows Phone\", \"capabilities\": { \"versioncontrol\": { \"sourceControlType\": \"Git\"}, \"processTemplate\": {  \"templateTypeId\": \"6b724908-ef14-45cf-84f8-768b5384da45\" }}}"

	var jsonStr = []byte(jsonFormat)

	var baseURL = "https://dev.azure.com/" + organization + "/_apis/projects?api-version=5.0"

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonStr))

	basic := "Basic " + patEncoded

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
