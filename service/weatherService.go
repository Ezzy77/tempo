package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ezzy77/tempo/models"
)

func GetCurrentWeather(location string, apiKey string) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, location)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather models.Weather

	json.Unmarshal(body, &weather)

	fmt.Printf("Weather in %s:\n", weather.Location.Name)
	fmt.Printf("Temperature: %.2f°C\n", weather.Current.TempC)
	fmt.Printf("Description: %s\n", weather.Current.Condition.Text)
}
