# go-lol-esports

**Golang application to get League of Legends Esports information (matches, leagues, etc.) using the Pandascore API and the Leaguepedia website.**

### Installation
You need to install Golang on your machine. If it's not installed, be sure to
[install Golang](https://golang.org/doc/install).

After this, you need to configure your Pandascore API key as an environment variable.
```sh
echo 'export PANDASCORE_KEY="<apiKey"' >> <terminalFile>
```

To get your `<apiKey>` go to (Pandascore
settings)[https://pandascore.co/settings] and you'll be able to see your API
key. As for the `<terminalFile`, it depends on your terminal. If you're using
bash, change it for `~/.bashrc`, as for zsh, you can change for `~/.zshenv`.

You can use a `config.json` if you want to add some configurations. The variable
names are trivial. Check the following structure:
```json
{
    "ignoreLeagues": ["CBLOL", "LCS"],
    "showResults": true 
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
-l, --league CODE SPLIT Prints the standings for a split of a league (e.g. go-lol-esports.exe -l LCS Spring)
-l, --league CODE SPLIT WEEK Prints the matches of a week of a split of a league (e.g. go-lol-esports.exe -l LCS Spring W3)

```

### Examples

#### Show today's matches

`$ go-lol-esports.exe`

![ScreenShot](/img/today_matches.jpg)

#### Show standings for a split of a league

`$ go-lol-esports.exe -l LCS Spring`

![ScreenShot](/img/league_standings.jpg)

#### Show results for a week of a split of a league

`$ go-lol-esports.exe -l LCS Spring W3`

![ScreenShot](/img/league_week_results.jpg)

### Contribute

Found a bug? Have a feature you'd like to see added or something you'd like to see improved? You can report it to me by [opening a new issue](https://github.com/aosousa/go-lol-esports/issues)!

### License

MIT © [André Sousa](https://github.com/aosousa)
