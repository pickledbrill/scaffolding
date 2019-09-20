package scaffolding

import (
	"testing"
)

func TestLoadConfig_LoadDefault_ShouldSuccess(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}
	templates := config.TemplateStructure
	findDefault := false
	for _, template := range templates {
		if template.Name == "default" {
			findDefault = true
		}
	}
	if !findDefault {
		t.Error("Failed to load default project template")
	}
}
