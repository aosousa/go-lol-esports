package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

/*Config struct contains all the necessary configurations for the app
to work correctly. */
type Config struct {
	APIKey        string   `json:"apiKey"`
	IgnoreLeagues []string `json:"ignoreLeagues"`
	ShowResults   bool     `json:"showResults"`
}

// CreateConfig adds information to a Config struct
func CreateConfig() Config {
	var config Config
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		utils.HandleError(err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		utils.HandleError(err)
	}

	if config.APIKey == "" {
		fmt.Println("ERROR: PandaScore API key missing.")
		os.Exit(1)
	}

	return config
}
