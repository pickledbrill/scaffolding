package request

import (
	"fmt"
	"net/http"
)

type HTTPRequest struct{}

// HTTPClient is the client to send http request.
var HTTPClient *http.Client

// GithubV3URL is the github version 3 endpoint url
const GithubV3URL string = "https://api.github.com"

// func AuthenticateUser() {
// 	requestURL := buildURL("authenticate")
// 	// a := http.NewRequest("GET", reque)
// }

// DispatchRequest starts a http request.
func DispatchRequest(request *http.Request) (interface{}, error) {
	token := ""
	request.Header.Add("Authentication", fmt.Sprintf("token %s", token))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	return nil, nil
}

// buildURL concats the http request url for different requests
func buildURL(action string) string {
	var urlPath string
	if exist := CheckURLExist(action); !exist {
		urlPath = ""
	} else {
		urlPath = GetMappedURL(action)
	}
	return fmt.Sprintf("%s%s", GithubV3URL, urlPath)
}
