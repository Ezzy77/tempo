package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Ezzy77/tempo/models"
)

func GetCurrentWeather(location string, apiKey string) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, location)
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

	city, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s\n", city.Name, city.Country, current.TempC, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		fmt.Printf(
			"%s - %.0fC, %.0f%% %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceRain,
			hour.Condition.Text)

	}

	// fmt.Printf("Weather in %s:\n", weather.Location.Name)
	// fmt.Printf("Temperature: %.2f°C\n", weather.Current.TempC)
	// fmt.Printf("Description: %s\n", weather.Current.Condition.Text)
}
