package main

import (
	"fmt"
	"os"
  "time"

	"CLIAPPS/data"
	"github.com/janeczku/go-spinner"
	"github.com/Delta456/box-cli-maker/v2"
)

type WeatherIcon struct {
	Code string
	Icon string
}

var weatherIcons = []WeatherIcon{
	{"01d", "ð"},
	{"01n", "ð"},
	{"02d", "ð¤"},
	{"02n", "ð¤"},
	{"03d", "âï¸"},
	{"03n", "âï¸"},
	{"04d", "âï¸"},
	{"04n", "âï¸"},
	{"09d", "ð§"},
	{"09n", "ð§"},
	{"10d", "ð¦"},
	{"10n", "ð§"},
	{"11d", "â"},
	{"11n", "â"},
	{"13d", "âï¸"},
	{"13n", "âï¸"},
	{"50d", "ð¬"},
	{"50n", "ð¬"},
}

func main() {
	
  Box := box.New(box.Config{Px: 2, Py: 5, Type: "Bold", TitlePos: "Top"})

  // Ask the user for the city location
	fmt.Print("Enter the city location: ")
	var city string
	fmt.Scan(&city)

	// Get the weather data
	data, err := data.GetWeatherData(city, "<your_API_key>")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the weather code
	weatherCode := data.Weather[0].Code

	// Get the appropriate icon
	var icon string
	for _, w := range weatherIcons {
		if w.Code == weatherCode {
			icon = w.Icon
			break
		}
	}
		s := spinner.StartNew("Processing...")
		s.SetSpeed(100 * time.Millisecond)
		// s.SetCharset([]string{"â£¾", "â£½", "â£»", "â¢¿", "â¡¿", "â£", "â£¯", "â£·"})

		s.SetCharset([]string{
			"â±â±â±â±â±",
			"â°â±â±â±â±",
			"â°â°â±â±â±",
			"â°â°â°â±â±",
			"â°â°â°â°â±",
			"â°â°â°â°â°",
			"â°â°â°â°â±",
			"â°â°â°â±â±",
			"â°â°â±â±â±",
			"â°â±â±â±â±",
			"â±â±â±â±â±",
    })

		time.Sleep(3 * time.Second)
		s.Stop()


weatherData := fmt.Sprintf("ðHere is your weather report..\n\n%s %s\nð Region: %s - %s\nð¡ï¸ Current Temperature: %.2fÂ°C\nð§ Precipitation: %.2f%%\nð¦ Humidity: %.2f%%\nð¨ Wind: %.2f km/h\n",
	icon, data.Weather[0].Description, city, data.Sys.Country, data.Main.Temp, data.Rain.OneHour, data.Main.Humidity, data.Wind.Speed)

Box.Print("Go-WeatherCLI", weatherData)

  }
