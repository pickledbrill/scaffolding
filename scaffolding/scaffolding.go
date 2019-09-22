package scaffolding

import (
	"errors"
	"strings"

	appErrors "github.com/pickledbrill/scaffolding/scaffolding/error"
)

// VerifyFlags verifies the valid status for each flag.
func VerifyFlags(args ...interface{}) error {
	template := args[0].(string)
	name := args[1].(string)

	if strings.TrimSpace(name) == "" {
		return errors.New(appErrors.InvalidRepositoryNameError)
	}
	if !checkTemplateExist(template) {
		return errors.New(appErrors.RepositoryTemplateNotFoundError)
	}
	return nil
}

// Process handles user validation and repository creation.
func Process(template string, name string, private bool) {
	// authenticate user
	err := AuthenticateUser()
	if err != nil {
		panic(err)
	}
	// create repository
	postBody := RepositoryPost{
		Name:    name,
		Private: private,
	}
	err = CreateRepository(postBody)
	if err != nil {
		panic(err)
	}
}

func checkTemplateExist(name string) bool {
	result := false
	for _, template := range GitConfig.TemplateStructure {
		if template.Name == name {
			result = true
			break
		}
	}
	return result
}
