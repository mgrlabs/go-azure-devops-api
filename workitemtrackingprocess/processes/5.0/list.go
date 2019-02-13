package processlist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiVersion = "5.0-preview.2"
var baseURL = "https://dev.azure.com/"
var apiURL = "/_apis/work/processes?api-version="

func listProcessTemplates() {

}
