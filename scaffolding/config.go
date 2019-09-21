package scaffolding

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config stores the config data stored in config.json file.
type Config struct {
	Access            AccessConfig    `json:"access"`
	TemplateStructure []RepoStructure `json:"repo_structure"`
}

// AccessConfig stores user's github access.
type AccessConfig struct {
	AccessToken string `json:"access_token"`
}

// RepoStructure stores a list of project structure.
type RepoStructure struct {
	Name  string   `json:"structure_name"`
	Files []string `json:"files"`
}

// GitConfig stores the configuration value.
var GitConfig Config

// InitializeConfig sets the configuration value to GitConfig variable.
func InitializeConfig() {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}
	GitConfig = *config
}

// LoadConfig reads the config.json file and load configuration.
func loadConfig() (*Config, error) {
	jsonFile, err := os.Open("../config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	return &config, nil
}
