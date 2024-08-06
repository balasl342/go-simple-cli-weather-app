package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// WeatherResponse defines the structure of the API response
type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

// Configuration struct
type Config struct {
	APIKey string `json:"api_key"`
}

func main() {
	city := flag.String("city", "London", "City to get weather for")
	flag.Parse()

	// Load configuration
	apiKey, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Fetch weather data
	weather, err := fetchWeather(*city, apiKey)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	// Print the weather data
	printWeather(weather)
}

func fetchWeather(city, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func printWeather(weather *WeatherResponse) {
	fmt.Printf("City: %s\n", weather.Name)
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d%%\n", weather.Main.Humidity)
	fmt.Printf("Weather: %s\n", weather.Weather[0].Description)
}

func loadConfig(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return "", err
	}

	return config.APIKey, nil
}
