package request

import (
	"fmt"
	"testing"
)

func TestBuildURL_ShouldSuccess_WithCorrectAction(t *testing.T) {
	// setup
	expected := fmt.Sprintf("%s%s", GithubV3URL, "/user")
	result := buildURL("authenticate")
	// validate
	if expected != result {
		t.Error("URL is not correct")
	}
}

func TestBuildURL_ShouldSuccess_WithWrongAction(t *testing.T) {
	// setup
	expected := fmt.Sprintf("%s%s", GithubV3URL, "")
	result := buildURL("test")
	// validate
	if expected != result {
		t.Error("URL is not correct")
	}
}
