package processlist

import (
	"io/ioutil"
	"log"
	"net/http"
)

var apiVersion = "5.0-preview.2"
var baseURI = "https://dev.azure.com/"
var apiPath = "/_apis/work/processes?api-version="

func listProcessTemplates(encodedPAT, azureDevopsOrg string) (processTemplates string) {

	requestURL := baseURI + azureDevopsOrg + apiPath + apiVersion
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("Authorization", "Basic "+encodedPAT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Close the response body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	processTemplates = string(body)
	return processTemplates
}
