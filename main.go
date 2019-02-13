package main

import "os"

func init() {
	// set up Config struct before performing any queries
	initConfig()
}

func main() {
	args := os.Args

	if len(args) > 1 {
		cmd := args[1]

		switch cmd {
		case "-l", "--league":
			handleLeagueOptions(args)
		case "-h", "--help":
			printHelp()
		case "-v", "--version":
			printVersion()
		}
	} else {
		getTodayMatches()
	}
}
