package models

import "fmt"

/*Match is the struct that represents a League of Legends Esports match.
 * Contains:
 * Winner struct - information about the winner of the match (used to get the winner of the match)
 * Opponents struct - detailed information about the teams facing each other (along with Results struct, it's used to get the score of the match)
 * Serie struct - detailed information about the league split the teams are competing in (along with League struct, used to get the league's name and split)
 * League struct - detailed information about the league the teams are competing in (along with Serie struct, used to get the league's name and split)
 * Results struct - information about the score of the match (along with Opponents struct, it' used to get the score of match)
 */
type Match struct {
	Winner    Winner     `json:"winner"`
	Opponents []Opponent `json:"opponents"`
	Name      string     `json:"name"`
	Serie     Serie      `json:"serie"`
	Results   []Result   `json:"results"`
	League    League     `json:"league"`
}

// Matches represents a slice of Match structs
type Matches []Match

// Serie is the struct that represents a League of Legends league split.
type Serie struct {
	Year   int    `json:"year"`
	Slug   string `json:"slug"`
	Season string `json:"season"`
	Name   string `json:"full_name"`
	Date   string `json:"begin_at"`
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

// Result is the struct that contains information about the score of a match.
type Result struct {
	TeamID int `json:"team_id"`
	Score  int `json:"score"`
}

// Error is the struct that contains information about a possible error in an HTTP request to the API
type Error struct {
	Error string `json:"error"`
}

// PrintMatches takes the information of each Match struct stored in a Matches slice and prints it out to the user
func (m Matches) PrintMatches() {
	for _, match := range m {
		leagueName := fmt.Sprintf("%s %s", match.League.Name, match.Serie.Name)
		teams := fmt.Sprintf("%s vs %s", match.Opponents[0].Team.Acronym, match.Opponents[1].Team.Acronym)
		formattedDate := fmt.Sprintf("%s %s", match.Serie.Date[0:10], match.Serie.Date[11:19])
		fmt.Printf("[%s] %s (%s)\n", formattedDate, teams, leagueName)
	}
}
