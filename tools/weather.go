package tools

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
)

// WeatherData represents the weather info (adjust fields as needed)
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name        string `json:"name"`
	Temperature float64
	Condition   string
}

type ClientWeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

func GetWeather(city string) (ClientWeatherResponse, error) {

	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid="+"your-api-key-here", encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return ClientWeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return ClientWeatherResponse{}, err
	}

	if len(weather.Weather) == 0 {
		return ClientWeatherResponse{}, fmt.Errorf("no weather data returned for city: %s", city)
	}

	celsius := weather.Main.Temp - 273.15
	roundedTemp := math.Round(celsius)

	clientResponse := ClientWeatherResponse{
		Temperature: roundedTemp, // Convert Kelvin to Celsius
		Condition:   weather.Weather[0].Description,
	}

	return clientResponse, nil
}
