# Premier League CLI App

This is a command-line interface (CLI) app built with Golang that displays information about the Premier League. It provides information about specific players, fixtures of games, and the current table positions.

## Installation

Clone this repository and then simply build the app with this command:
```Golang
go build .

```

## Usage
The app provides three commands: find player, find fixture, and find table.

To find information about a specific player, run the following command:

```Golang
premierleagueCLI find player <name_of_the_player>

```

To find information about fixtures, run the following command:

```Golang
premierleagueCLI find fixture

```
To find information about table positions, run the following command:

```Golang
premierleagueCLI find table

```

## Libraries used
This app uses the following third-party libraries:

- [goquery](https://github.com/PuerkitoBio/goquery) a package that brings a syntax and a set of features similar to jQuery to the Go language.
- [simpletable](https://github.com/alexeyco/simpletable) a package that allows to easily create tables in Go.
- [cobra](https://github.com/spf13/cobra) a package that helps in creating powerful modern CLI interfaces.



## Credits
The Premier League API used by this app is provided by [tarun7r/Premier-League-API](https://github.com/tarun7r/Premier-League-API). The table positions information is webscraped manually from [onefootball](https://onefootball.com/en/competition/premier-league-9/table).
