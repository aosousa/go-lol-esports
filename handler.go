package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"./models"
	"./utils"

	"github.com/anaskhan96/soup"
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
		// no leagues to ignore - print today's matches as they are
		if len(config.IgnoreLeagues) == 0 {
			matches.PrintMatches()
		} else {
			finalMatches := removeIgnoredMatches(config.IgnoreLeagues, matches)
			finalMatches.PrintMatches()
		}
	} else {
		fmt.Println("No matches today.")
	}
}

/* Removes matches from ignored leagues from today's matches results
using the FindInSlice method above */
func removeIgnoredMatches(ignoreLeagues []string, matches models.Matches) models.Matches {
	var finalMatches models.Matches

	for _, match := range matches {
		if !utils.FindInSlice(match.League.Name, ignoreLeagues) {
			finalMatches = append(finalMatches, match)
		}
	}

	return finalMatches
}

/* Handles a request to lookup information related to a league (week results, standings)
 * Receives:
 * args ([]string) - Arguments passed in the terminal by the user
 */
func handleLeagueOptions(args []string) {
	var queryURL string

	lastArg := args[len(args)-1]
	weekRegex, _ := regexp.Compile("(?i)W[0-9]+")
	weekRegexArg := weekRegex.FindStringSubmatch(lastArg)

	// if week argument was not sent in the command line, perform query for league standings
	// if it was, perform query for a league's week results
	if len(weekRegexArg) == 0 {
		leagueCode := strings.ToUpper(args[len(args)-2])
		split := strings.Title(lastArg)

		queryURL = fmt.Sprintf("%s%s/%d_Season/%s_Season", baseLeaguepediaURL, leagueCode, time.Now().Year(), split)
		resp, err := soup.Get(queryURL)
		if err != nil {
			utils.HandleError(err)
		}

		printLeagueTable(leagueCode, split, resp)
	}
}

/* Print a league's standings
 * Receives:
 * leagueCode (string) - Code of the league (LCS, LEC, etc.)
 * split (string) - Split of the league season (Spring, Summer)
 * document (string) - Leaguepedia page HTML document
 */
func printLeagueTable(leagueCode string, split string, document string) {
	currentYear := time.Now().Year()
	fmt.Printf("%s %d %s Standings\n", leagueCode, currentYear, split)

	doc := soup.HTMLParse(document)
	standings := doc.Find("table", "class", "wikitable2")
	teams := standings.FindAll("tr", "class", "teamhighlight")
	for _, team := range teams {
		td := team.FindAll("td")
		teamSpan := td[1].Find("span", "class", "teamname")
		teamName := teamSpan.FindAll("a")[0].Text()
		teamResults := strings.Replace(td[2].Text(), " ", "", -1)
		standing := fmt.Sprintf("%s. %s [%s]", td[0].Text(), teamName, teamResults)
		fmt.Println(standing)
	}
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Println("Go LoL Esports (version " + version + ")")
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help 	  Prints the list of available commands")
	fmt.Println("* -v | --version Prints the version of the application")
	fmt.Println("\n* -l | --league `league code` `split` Prints the standings for `split` of `league code` (e.g. go-lol-esports.exe -l LEC Spring)")
	fmt.Println("If no arguments are sent (e.g. go-lol-esports.exe), the application prints the current day's matches.")
}

// Prints the current version of the application
func printVersion() {
	fmt.Println("Go LoL Esports version " + version)
}
