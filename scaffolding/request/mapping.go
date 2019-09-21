package request

// RequestURLMap stores the url mapping for different requests.
var RequestURLMap = map[string]string{
	"authenticate": "/user",
}

// CheckURLExist checks whether the http request action has a mapping URL.
func CheckURLExist(action string) bool {
	_, ok := RequestURLMap[action]
	return ok
}

// GetMappedURL returns the mapped URL.
func GetMappedURL(action string) string {
	return RequestURLMap[action]
}
