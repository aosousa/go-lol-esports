package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aosousa/go-lol-esports/utils"
)

/*Match is the struct that represents a League of Legends Esports match.
 * Contains:
 * Winner struct - information about the winner of the match (used to get the winner of the match)
 * Opponents struct - detailed information about the teams facing each other (along with Results struct, it's used to get the score of the match)
 * Serie struct - detailed information about the league split the teams are competing in (along with League struct, used to get the league's name and split)
 * League struct - detailed information about the league the teams are competing in (along with Serie struct, used to get the league's name and split)
 * Results struct - information about the score of the match (along with Opponents struct, it' used to get the score of match)
 */
type Match struct {
	Winner    Winner    `json:"winner"`
	Opponents Opponents `json:"opponents"`
	Name      string    `json:"name"`
	Serie     Serie     `json:"serie"`
	Results   Results   `json:"results"`
	League    League    `json:"league"`
	Date      string    `json:"begin_at"`
	Status    string    `json:"status"`
}

// Matches represents a slice of Match structs
type Matches []Match

// Serie is the struct that represents a League of Legends league split.
type Serie struct {
	Year   int    `json:"year"`
	Slug   string `json:"slug"`
	Season string `json:"season"`
	Name   string `json:"full_name"`
}

// League is the struct that represents a League of Legends league.
type League struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Winner is the struct that contains information about the team who won a match.
type Winner struct {
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}

// Opponent is the struct that contains information about a team playing in a match.
type Opponent struct {
	Team struct {
		Name    string `json:"name"`
		ID      int    `json:"id"`
		Acronym string `json:"acronym"`
	} `json:"opponent"`
}

// Opponents represents a slice of Opponent structs
type Opponents []Opponent

// Result is the struct that contains information about the score of a match.
type Result struct {
	TeamID int `json:"team_id"`
	Score  int `json:"score"`
}

// Results represents a slice of Result structs
type Results []Result

// Error is the struct that contains information about a possible error in an HTTP request to the API
type Error struct {
	Error string `json:"error"`
}

/*PrintMatches takes the information of each Match struct stored in a Matches slice and prints it out to the user.
 * Receives:
 * showResults (bool) - Whether or not to show results if they are available
 * liveMatches (bool) - Whether or not that set of matches is from the /running endpoints
 */
func (m Matches) PrintMatches(showResults bool) {
	for _, match := range m {
		var teams, formattedDate string

		leagueName := fmt.Sprintf("%s %s", match.League.Name, match.Serie.Name)

		// match has a score and showResults is true - show the scores
		if match.hasScores() && showResults {
			var firstTeamScore, secondTeamScore int

			if match.Opponents[0].Team.Acronym == match.Winner.Acronym {
				firstTeamScore, secondTeamScore = match.Results[0].Score, match.Results[1].Score
			} else {
				firstTeamScore, secondTeamScore = match.Results[1].Score, match.Results[0].Score
			}

			teams = fmt.Sprintf("%s %d - %d %s", match.Opponents[0].Team.Acronym, firstTeamScore, secondTeamScore, match.Opponents[1].Team.Acronym)
		} else {
			teams = fmt.Sprintf("%s vs %s", match.Opponents[0].Team.Acronym, match.Opponents[1].Team.Acronym)
		}

		// live match is running but does not have a winner yet - add the LIVE tag
		if match.Status == "running" {
			formattedDate = fmt.Sprintf("%s %s LIVE", match.Date[0:10], match.Date[11:19])
		} else {
			formattedDate = fmt.Sprintf("%s %s", match.Date[0:10], match.Date[11:19])
		}

		fmt.Printf("[%s] %s (%s)\n", formattedDate, teams, leagueName)
	}
}

/*GetMatches gets past or upcoming matches from the current date and returns the Matches struct
 * Receives:
 * url (string) - Endpoint URL
 *
 * Returns: Matches struct
 */
func (m Matches) GetMatches(url string, sort bool) Matches {
	res, err := http.Get(url)
	if err != nil {
		utils.HandleError(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		utils.HandleError(err)
	}

	json.Unmarshal(content, &m)

	if sort {
		return m.sortMatches()
	}

	return m
}

/* Sorts matches by date by looping through the original slice in reverse
before sending them to the handler.*/
func (m Matches) sortMatches() Matches {
	var newMatchSlice Matches

	for i := len(m) - 1; i >= 0; i-- {
		newMatchSlice = append(newMatchSlice, m[i])
	}

	return newMatchSlice
}

/* Checks if a match already has scores to show */
func (m Match) hasScores() bool {
	if m.Results[0].Score == 0 && m.Results[1].Score == 0 {
		return false
	}

	return true
}
