package scaffolding

import (
	"testing"

	appErrors "github.com/pickledbrill/scaffolding/scaffolding/error"
)

func TestCheckTemplateExist_ShouldSuccess(t *testing.T) {
	// setup
	GitConfig.TemplateStructure = make([]RepoStructure, 1)
	GitConfig.TemplateStructure[0] = RepoStructure{
		Name: "default",
	}
	expectedResult := true
	result := checkTemplateExist("default")
	// validate
	if expectedResult != result {
		t.Error("Result doesn't match")
	}
}

func TestCheckTemplateExist_ShouldFail(t *testing.T) {
	// setup
	GitConfig.TemplateStructure = make([]RepoStructure, 1)
	GitConfig.TemplateStructure[0] = RepoStructure{
		Name: "default",
	}
	expectedResult := false
	result := checkTemplateExist("test")
	// validate
	if expectedResult != result {
		t.Error("Result doesn't match")
	}
}

func TestVerifyFlags_ShouldSuccess(t *testing.T) {
	// setup
	GitConfig.TemplateStructure = make([]RepoStructure, 1)
	GitConfig.TemplateStructure[0] = RepoStructure{
		Name: "default",
	}
	template := "default"
	name := "www"
	err := VerifyFlags(template, name)
	// validate
	if err != nil {
		t.Error(err)
	}
}

func TestVerifyFlags_ShouldFail_WithEmptyRepositoryName(t *testing.T) {
	// setup
	template := "default"
	name := ""
	err := VerifyFlags(template, name)
	// validate
	if err == nil {
		t.Error("Should return error")
	}
	if err != nil && err.Error() != appErrors.InvalidRepositoryNameError {
		t.Error("Expected error doesn't match")
	}
}

func TestVerifyFlags_ShouldFail_WithInvalidRepositoryTemplate(t *testing.T) {
	// setup
	GitConfig.TemplateStructure = make([]RepoStructure, 1)
	GitConfig.TemplateStructure[0] = RepoStructure{
		Name: "default",
	}
	template := "test"
	name := "test_repo"
	err := VerifyFlags(template, name)
	// validate
	if err == nil {
		t.Error("Should return error")
	}
	if err != nil && err.Error() != appErrors.RepositoryTemplateNotFoundError {
		t.Error("Expected error doesn't match")
	}
}
