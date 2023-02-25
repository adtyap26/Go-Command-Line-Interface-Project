package config

import (
  "encoding/json"
  "fmt"
  "os"
)

// Config to store API key
type Configuration struct {
  APIKey string `json:"apiKey"`
}

// Now we load the that config with function
func LoadConfig(file string) Configuration  {
  var config Configuration
  configFile, err := os.Open(file)
  defer configFile.Close()
  if err != nil {
    fmt.Println(err.Error())
  }
  jsonParser := json.NewDecoder(configFile)
  jsonParser.Decode(&config)
  return config
}
