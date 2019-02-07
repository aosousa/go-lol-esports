package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lol-esports-calendar/models"
	"lol-esports-calendar/utils"
	"net/http"
	"time"
)

const (
	basePandascoreURL  = "https://api.pandascore.co/lol/"
	baseLeaguepediaURL = "https://lol.gamepedia.com/"
	version            = "1.0.0"
)

var config models.Config

// Read configuration file and parse that information into a Config struct
func initConfig() {
	config = models.CreateConfig()
}

// Gets today's professional LoL matches through the Pandascore API
func getTodayMatches() {
	var (
		queryURL    string
		matches     models.Matches
		currentDate string
	)

	currentDate = time.Now().Local().Format("2006-01-02")
	queryURL = basePandascoreURL + "matches/upcoming?token=" + config.APIKey + "&filter[begin_at]=" + currentDate

	res, err := http.Get(queryURL)
	if err != nil {
		utils.HandleError(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		utils.HandleError(err)
	}

	json.Unmarshal(content, &matches)

	if len(matches) > 0 {
		matches.PrintMatches()
	} else {
		fmt.Println("No matches today.")
	}
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Println("LoL Esports Calendar (version " + version + ")")
	fmt.Println("Available commands:")
}

// Prints the current version of the application
func printVersion() {
	fmt.Println("Version " + version)
}
