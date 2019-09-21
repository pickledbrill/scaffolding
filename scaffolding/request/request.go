package request

import (
	"fmt"
	"net/http"
)

// HTTPClient is the client to send http request.
var HTTPClient *http.Client = &http.Client{}

// DispatchRequest starts a http request.
func DispatchRequest(request *http.Request, authToken string) (*http.Response, error) {
	request.Header.Add("Authorization", fmt.Sprintf("token %s", authToken))
	request.Header.Add("Accept", "application/vnd.github.v3+json")
	return HTTPClient.Do(request)
}
