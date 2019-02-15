package processlist

import (
	"bytes"
	"net/http"
)

// https://docs.microsoft.com/en-us/rest/api/azure/devops/processes/processes/list?view=azure-devops-rest-5.0

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

var apiVersion = "5.0-preview.2"
var baseURL = "https://dev.azure.com/"
var apiURL = "/_apis/work/processes?api-version="

func listProcessTemplates(encodedPAT, organization string) (processList map[string]string) {

	callURL := baseURL + organization + apiURL + apiVersion
	req, err := http.NewRequest("POST", callURL, bytes.NewBuffer(jsonStr))
	basic := "Basic " + encodedPAT

	req.Header.Set("Authorization", basic)
	return
}
