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
	workDir := args[2].(string)

	if strings.TrimSpace(name) == "" {
		return errors.New(appErrors.InvalidRepositoryNameError)
	}
	if !checkTemplateExist(template) {
		return errors.New(appErrors.RepositoryTemplateNotFoundError)
	}
	if strings.TrimSpace(workDir) == "" {
		return errors.New(appErrors.InvalidDirectoryError)
	}
	return nil
}

// Process handles user validation and repository creation.
func Process(template string, name string, private bool, workDir string) {
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
	returnResult, err := CreateRepository(postBody)
	if err != nil {
		panic(err)
	}

	templatePtr := getTemplate(template)
	// setup repository
	ok, err := SetupRepository(workDir, templatePtr.Files)
	if !ok {
		panic(err)
	}
	GitInitRepository(workDir, *returnResult)
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

func getTemplate(name string) *RepoStructure {
	structure := &RepoStructure{}
	for _, template := range GitConfig.TemplateStructure {
		if template.Name == name {
			structure = &template
			break
		}
	}
	return structure
}
