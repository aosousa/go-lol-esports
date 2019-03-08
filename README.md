# go-lol-esports

**Golang application to get League of Legends Esports information (matches, leagues, etc.) using the Pandascore API and the Leaguepedia website.**

### Installation
Download the executable in the [Releases tab](https://github.com/aosousa/go-lol-esports/releases).

This tool requires a `config.json` file present in the same directory as the executable, with the following structure:
```json
{
    "apiKey": "<your-Pandascore-API-key>",
    "ignoreLeagues": "<Array of leagues to ignore> (i.e. ['LCS', 'VCS'])",
    "showResults": true "(Whether or not to show match results if they are available)"
}
```

Available league codes (for the ignoreLeagues configuration):

| League Code | Description |
| --- | --- |
| LCS | LoL Championship Series (NA) |
| LEC | LoL European Championship | 
| LCK | LoL Champions Korea | 
| Challenger Korea | LoL Challengers Korea |
| LPL | LoL Pro League (China) |
| LMS | League Master Series (Taiwan) |
| CBLOL | Campeonato Brasileiro LoL (Brazil) |
| LCL | LoL Russia League |
| LJL | LoL Japan League |
| LLA | Latin America League |
| OPL | Oceanic Pro League |
| LST | LoL SEA Tour |
| TCL | Turkish Champions League |
| VCS | Vietnam Championship Series |
| LVP SLO | SuperLiga Orange (Spain) |

### Usage
```
go-lol-esports.exe [-l | --league | -h | --help | -v | --version]
```

### Options

```
-h, --help Prints the list of available commands
-v, --version Prints the version of the application
-l, --league CODE SPLIT Prints the standings for a split of a league (e.g. go-lol-esports.exe -l LEC Spring)
-l, --league CODE SPLIT WEEK Prints the matches of a week of a split of a league (e.g. go-lol-esports.exe -l LEC Spring W3)

```

### Examples

### Contribute

Found a bug? Have a feature you'd like to see added or something you'd like to see improved? You can report it to me by [opening a new issue](https://github.com/aosousa/go-lol-esports/issues)!

### License

MIT © [André Sousa](https://github.com/aosousa)