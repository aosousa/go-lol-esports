package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	utils "github.com/aosousa/golang-utils"
)

/*Config struct contains all the necessary configurations for the app
to work correctly. */
type Config struct {
	APIKey        string   `json:"apiKey"`
	IgnoreLeagues []string `json:"ignoreLeagues"`
	ShowResults   bool     `json:"showResults"`
}

func checkErr(err error) {
	if err != nil {
		utils.HandleError(err)
	}
}

// CreateConfig adds information to a Config struct
func CreateConfig() Config {
	var config Config

	if _, err := os.Stat("./config.json"); err == nil {
		jsonFile, err := ioutil.ReadFile("./config.json")
		checkErr(err)
		err = json.Unmarshal(jsonFile, &config)
		checkErr(err)
	}

	config.APIKey = os.Getenv("PANDASCORE_KEY")

	if config.APIKey == "" {
		fmt.Println("ERROR: PandaScore API key missing.")
		os.Exit(1)
	}

	return config
}
