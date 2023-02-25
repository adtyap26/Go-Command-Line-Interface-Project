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
	{"01d", "🌞"},
	{"01n", "🌛"},
	{"02d", "🌤"},
	{"02n", "🌤"},
	{"03d", "☁️"},
	{"03n", "☁️"},
	{"04d", "☁️"},
	{"04n", "☁️"},
	{"09d", "🌧"},
	{"09n", "🌧"},
	{"10d", "🌦"},
	{"10n", "🌧"},
	{"11d", "⛈"},
	{"11n", "⛈"},
	{"13d", "❄️"},
	{"13n", "❄️"},
	{"50d", "🌬"},
	{"50n", "🌬"},
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
		// s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})

		s.SetCharset([]string{
			"▱▱▱▱▱",
			"▰▱▱▱▱",
			"▰▰▱▱▱",
			"▰▰▰▱▱",
			"▰▰▰▰▱",
			"▰▰▰▰▰",
			"▰▰▰▰▱",
			"▰▰▰▱▱",
			"▰▰▱▱▱",
			"▰▱▱▱▱",
			"▱▱▱▱▱",
    })

		time.Sleep(3 * time.Second)
		s.Stop()


weatherData := fmt.Sprintf("🌍Here is your weather report..\n\n%s %s\n📍 Region: %s - %s\n🌡️ Current Temperature: %.2f°C\n💧 Precipitation: %.2f%%\n💦 Humidity: %.2f%%\n💨 Wind: %.2f km/h\n",
	icon, data.Weather[0].Description, city, data.Sys.Country, data.Main.Temp, data.Rain.OneHour, data.Main.Humidity, data.Wind.Speed)

Box.Print("Go-WeatherCLI", weatherData)

  }
