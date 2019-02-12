package main

import "os"

func main() {
	// set up Config struct
	initConfig()

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
