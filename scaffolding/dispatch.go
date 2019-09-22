package scaffolding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	appErrors "github.com/pickledbrill/scaffolding/scaffolding/error"
	requestDispatch "github.com/pickledbrill/scaffolding/scaffolding/request"
)

// GithubV3URL is the github version 3 endpoint url
const GithubV3URL string = "https://api.github.com"

// AuthenticateUser authenticates the user's identity via sending request to
// Github v3 API.
func AuthenticateUser() error {
	requestURL := buildURL("authenticate")
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := requestDispatch.DispatchRequest(req, GitConfig.Access.AccessToken)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		return errors.New(appErrors.FailedUserAuthenticationError)
	}
	return nil
}

// CreateRepository sends the create repository http request to Github v3 API.
func CreateRepository(post RepositoryPost) error {
	requestURL := buildURL("createRepo")
	jsonValue, _ := json.Marshal(post)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonValue))
	resp, err := requestDispatch.DispatchRequest(req, GitConfig.Access.AccessToken)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 201 {
		return errors.New(appErrors.FailedRepositoryCreationError)
	}
	return nil
}

// buildURL concats the http request url for different requests
func buildURL(action string) string {
	var urlPath string
	if exist := requestDispatch.CheckURLExist(action); !exist {
		urlPath = ""
	} else {
		urlPath = requestDispatch.GetMappedURL(action)
	}
	return fmt.Sprintf("%s%s", GithubV3URL, urlPath)
}
