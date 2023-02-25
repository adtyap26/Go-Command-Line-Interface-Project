package data

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type WeatherData struct {
    Weather []struct {
        Description string `json:"description"`
        Code        string `json:"icon"`
    } `json:"weather"`
    Main struct {
        Temp     float64 `json:"temp"`
        Humidity float64 `json:"humidity"`
    } `json:"main"`
    Wind struct {
        Speed float64 `json:"speed"`
    } `json:"wind"`
    Rain struct {
        OneHour float64 `json:"1h"`
    } `json:"rain,omitempty"`
    Sys struct {
        Country string `json:"country"`
    } `json:"sys"`
}

type WeatherIcon struct {
    Code string
    Icon string
}

func GetWeatherData(city string, apiKey string) (*WeatherData, error) {
    // Construct the API url
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

    // Make the HTTP request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Unmarshal the JSON data into a struct
    var data WeatherData
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }

    return &data, nil
}
  
