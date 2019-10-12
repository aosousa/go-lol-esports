# go-lol-esports

**Golang application to get League of Legends Esports information (matches, leagues, etc.) using the Pandascore API and the Leaguepedia website.**

### Installation
You need to install Golang on your machine. If it's not installed, be sure to
[install Golang](https://golang.org/doc/install).

Alternatively, you can download the executable in the [Releases tab](https://github.com/aosousa/go-lol-esports/releases).

After this, you need to configure your Pandascore API key as an environment variable.
```sh
echo 'export PANDASCORE_KEY="<apiKey>"' >> <terminalFile>
```

To get your `<apiKey>` go to [Pandascore
settings](https://pandascore.co/settings) and you'll be able to see your API
key. As for the `<terminalFile>`, it depends on your terminal. If you're using
bash, change it for `~/.bashrc`, as for zsh, you can change for `~/.zshenv`.

You can use a `config.json` if you want to add some configurations by following this structure:
```json
{
    "ignoreLeagues": ["CBLOL", "LCS"],
    "showResults": true 
}
```

Available league codes and splits:

| League Code | Description | Splits |
| --- | --- | --- |
| LCS | LoL Championship Series (NA) | Spring, Summer |
| LEC | LoL European Championship | Spring, Summer |
| LCK | LoL Champions Korea | Spring, Summer |
| Challenger Korea | LoL Challengers Korea | |
| LPL | LoL Pro League (China) | Spring, Summer |
| LMS | League Master Series (Taiwan) | Spring, Summer |
| CBLOL | Campeonato Brasileiro LoL (Brazil) | Winter, Summer |
| LCL | LoL Russia League | Spring, Summer |
| LJL | LoL Japan League | Spring, Summer |
| LLA | Latin America League | Opening, Closing |
| OPL | Oceanic Pro League | Split_1, Split_2 |
| LST | LoL SEA Tour | Spring, Summer |
| TCL | Turkish Champions League | Winter, Summer |
| VCS | Vietnam Championship Series | Spring, Summer |
| LVP SLO | SuperLiga Orange (Spain) | |

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
