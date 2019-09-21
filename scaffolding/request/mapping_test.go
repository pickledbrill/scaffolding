package request

import "testing"

func TestRequestURLMap_ShouldContainCorrectMapping(t *testing.T) {
	// setup
	action := "authenticate"
	expectedURL := "/user"
	val, ok := RequestURLMap[action]
	// validate
	if !ok {
		t.Error("Key should exist in the mapping model")
	}
	if val != expectedURL {
		t.Error("Expected URL doesn't match")
	}
}

func TestCheckURLExist_ShouldSuccess(t *testing.T) {
	// setup
	action := "authenticate"
	result := CheckURLExist(action)
	// validate
	if !result {
		t.Error("Key should exist in the mapping model")
	}
}

func TestGetMappedURL_ShouldSuccess(t *testing.T) {
	// setup
	action := "authenticate"
	expectedURL := "/user"
	result := GetMappedURL(action)
	// validate
	if result != expectedURL {
		t.Error("Expected URL doesn't match")
	}
}
