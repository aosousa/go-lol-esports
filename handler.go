package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/aosousa/go-lol-esports/models"
	"github.com/aosousa/go-lol-esports/utils"
	ut "github.com/aosousa/golang-utils"
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
		pastMatches, liveMatches, upcomingMatches, matches           models.Matches
		currentDate, pastQueryURL, runningQueryURL, upcomingQueryURL string
	)

	currentDate = time.Now().Local().Format("2006-01-02")
	pastQueryURL = basePandascoreURL + "matches/past?token=" + config.APIKey + "&filter[begin_at]=" + currentDate
	runningQueryURL = basePandascoreURL + "matches/running?token=" + config.APIKey
	upcomingQueryURL = basePandascoreURL + "matches/upcoming?token=" + config.APIKey + "&filter[begin_at]=" + currentDate

	pastMatches = pastMatches.GetMatches(pastQueryURL, true)
	liveMatches = liveMatches.GetMatches(runningQueryURL, true)
	upcomingMatches = upcomingMatches.GetMatches(upcomingQueryURL, false)
	matches = append(pastMatches, liveMatches...)
	matches = append(liveMatches, upcomingMatches...)

	if len(matches) > 0 {
		// no leagues to ignore - print today's matches as they are
		if len(config.IgnoreLeagues) == 0 {
			matches.PrintMatches(config.ShowResults)
		} else {
			finalMatches := removeIgnoredMatches(config.IgnoreLeagues, matches)
			finalMatches.PrintMatches(config.ShowResults)
		}
	} else {
		fmt.Println("No matches today.")
	}
}

// Removes matches from ignored leagues from today's matches results
func removeIgnoredMatches(ignoreLeagues []string, matches models.Matches) models.Matches {
	var finalMatches models.Matches

	for _, match := range matches {
		if !ut.FindStringInSlice(match.League.Name, ignoreLeagues) {
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
	var (
		queryURL, leagueCode, split, week string
		currentYear                       int
	)

	currentYear = time.Now().Year()

	lastArg := args[len(args)-1]
	weekRegex, _ := regexp.Compile("(?i)W[0-9]+")
	weekRegexArg := weekRegex.FindStringSubmatch(lastArg)

	// if week argument was not sent in the command line, perform query for league standings
	// if it was, perform query for a league's week results
	if len(weekRegexArg) == 0 {
		leagueCode = strings.ToUpper(args[len(args)-2])
		split = strings.Title(lastArg)

		queryURL = fmt.Sprintf("%s%s/%d_Season/%s_Season", baseLeaguepediaURL, leagueCode, time.Now().Year(), split)
		resp, err := soup.Get(queryURL)
		if err != nil {
			utils.HandleError(err)
		}

		printLeagueTable(leagueCode, split, currentYear, resp)
	} else {
		leagueCode = strings.ToUpper(args[len(args)-3])
		split = strings.Title(args[len(args)-2])
		week = weekRegexArg[0][1:]

		queryURL = fmt.Sprintf("%s%s/%d_Season/%s_Season", baseLeaguepediaURL, leagueCode, time.Now().Year(), split)
		resp, err := soup.Get(queryURL)
		if err != nil {
			utils.HandleError(err)
		}

		printWeekResults(leagueCode, split, week, currentYear, resp)
	}
}

/* Print a league's standings
 * Receives:
 * leagueCode (string) - Code of the league (LCS, LEC, etc.)
 * split (string) - Split of the league season (Spring, Summer)
 * year (int) - Current year
 * document (string) - Leaguepedia page HTML document
 */
func printLeagueTable(leagueCode string, split string, year int, document string) {
	fmt.Printf("%s %d %s Standings\n", leagueCode, year, split)

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

/* Print the matches of a week in a league
 * Receives:
 * leagueCode (string) - Code of the league (LCS, LEC, etc.)
 * split (string) - Split of the league season (Spring, Summer)
 * week (string) - Week number
 * year (int) - Current year
 * document (string) - Leaguepedia page HTML document
 */
func printWeekResults(leagueCode, split, week string, year int, document string) {
	fmt.Printf("%s %d %s Week %s Matches\n\n", leagueCode, year, split, week)

	className := fmt.Sprintf("matches-week%s", week)
	doc := soup.HTMLParse(document)
	matches := doc.FindAll("tr", "class", className)

	for _, match := range matches {
		// only go through the td elements that have information about a match
		if date, ok := match.Attrs()["data-date"]; ok {
			firstTeam := match.Find("td", "class", "matchlist-team1").Children()[0].Children()[0].Text()
			secondTeam := match.Find("td", "class", "matchlist-team2").Children()[0].Children()[1].Text()
			hour := strings.Replace(match.Find("td", "class", "matchlist-time-cell").Children()[1].Children()[0].Children()[0].Text()[10:], ",", ":", -1)
			scores := match.FindAll("td", "class", "matchlist-score")

			// game has a score - show it
			if len(scores) > 0 {
				firstTeamScore := scores[0].Text()
				secondTeamScore := scores[1].Text()

				fmt.Printf("[%s %s] %s %s - %s %s\n", date, hour, firstTeam, firstTeamScore, secondTeamScore, secondTeam)
			} else {
				fmt.Printf("[%s %s] %s vs %s\n", date, hour, firstTeam, secondTeam)
			}
		}
	}
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Printf("Go LoL Esports (version %s)\n", version)
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help\t Prints the list of available commands")
	fmt.Println("* -v | --version Prints the version of the application")
	fmt.Println("\n* -l | --league `league code` `split`\t     Prints the standings for `split` of `league code` (e.g. go-lol-esports.exe -l LEC Spring)")
	fmt.Println("* -l | --league `league code` `split` `week` Prints the matches of `week` of `split` of `league code` (e.g. go-lol-esports.exe -l LEC Spring W3)")
	fmt.Println("If no arguments are sent (e.g. go-lol-esports.exe), the application prints the current day's matches.")
}

// Prints the current version of the application
func printVersion() {
	fmt.Printf("Go LoL Esports version %s\n", version)
}
