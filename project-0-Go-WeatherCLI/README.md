# GO-weatherCLI

This is a command-line interface (CLI) weather application written in Go. The app allows you to check the current weather for a given city location.

## How to use

Run the command 

```
go run main.go
```

Or build it with

```
go build
```

The program will prompt you to enter the city location. Type in the name of the city and press enter.

The program will display the current weather information for the city, including the weather icon, description, region, current temperature, precipitation, humidity, and wind speed.

## Libraries used

- fmt: for input/output operations
- os: for interacting with the operating system
- time: for time-related operations
- github.com/janeczku/go-spinner: for creating a loading spinner
- github.com/Delta456/box-cli-maker/v2: wrapping the result with box
- encoding/json: for working with JSON data
- io/ioutil: for reading the response body
- net/http: for making HTTP requests

## Note
The weather data is provided by OpenWeatherMap. You need to sign up to get an API key from OpenWeatherMap, and replace the api key in data.go file with your own.
