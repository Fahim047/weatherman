package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(city string, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("city not found")
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("API call failed with status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}
