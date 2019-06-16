package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aosousa/go-lol-esports/utils"
)

// Match is the struct that represents a League of Legends eSports match.
type Match struct {
	Winner    Winner    `json:"winner"` // Information about the winner of the match
	Opponents Opponents `json:"opponents"` // Detailed information about the teams facing each other (along with Results struct, it's used to get the score of the match)
	Name      string    `json:"name"` // Match opponents in a string (i.e. C9 vs CLG)
	Serie     Serie     `json:"serie"` // Detailed information about the league split the teams are competing in (along with Serie struct, it's used to get the league's name and split)
	Results   Results   `json:"results"` // Information about the score of the match
	League    League    `json:"league"` // Detailed information about the league the teams are competing in (along with the Serie struct, it's used to get the league's name and split)
	Date      string    `json:"begin_at"` // Date/time of match kick off
	Status    string    `json:"status"` // Status of the match (running, finished)
}

// Matches represents a slice of Match structs
type Matches []Match

// Serie is the struct that represents a League of Legends league split.
type Serie struct {
	Name   string `json:"full_name"` // Name of the League of Legends league split
}

// League is the struct that represents a League of Legends league.
type League struct {
	Name string `json:"name"` // Name of the League of Legends league
}

// Winner is the struct that contains information about the team who won a match.
type Winner struct {
	Acronym string `json:"acronym"` // Acronym of the winner of the match
}

// Opponent is the struct that contains information about a team playing in a match.
type Opponent struct {
	Team struct {
		Name    string `json:"name"` // Name of the team
		Acronym string `json:"acronym"` // Acronym of the name of the team
	} `json:"opponent"`
}

// Opponents represents a slice of Opponent structs
type Opponents []Opponent

// Result is the struct that contains information about the score of a match.
type Result struct {
	Score  int `json:"score"` // Number of games won by a team in a match
}

// Results represents a slice of Result structs
type Results []Result

/*PrintMatches takes the information of each Match struct stored in a Matches slice and prints it out to the user.
 * Receives:
 * showResults (bool) - Whether or not to show results if they are available
 */
func (m Matches) PrintMatches(showResults bool) {
	// print finished matches
	finishedMatches := m.filterMatchesByStatus("finished")
	fmt.Println("Finished matches:")
	finishedMatches.printSection(showResults)

	// print live matches
	liveMatches := m.filterMatchesByStatus("running")
	fmt.Println("\nLive matches:")
	liveMatches.printSection(showResults)

	// print upcomming matches
	upcomingMatches := m.filterMatchesByStatus("not_started")
	fmt.Println("\nUpcoming matches:")
	upcomingMatches.printSection(showResults)
}

/* Prints a section of matches (finished, running, upcoming)
 * Receives:
 * showResults (bool) - Whether ot not to show results if they are available
 */
func (m Matches) printSection(showResults bool) {
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

/* Filter matches by their status (not started, finished, running) */
func (m Matches) filterMatchesByStatus(status string) Matches {
	var filteredMatches Matches

	for _, match := range m {
		if match.Status == status {
			filteredMatches = append(filteredMatches, match)
		}
	}

	return filteredMatches
}
