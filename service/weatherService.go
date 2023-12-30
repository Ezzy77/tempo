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
}

func GetForecastDays(location string, apiKey string, numbDays string) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%s", apiKey, location, numbDays)
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

	days := weather.Forecast.ForecastDay

	fmt.Print("\n\n")
	fmt.Println("Date        Min_Temp  Max_Temp    Avr_Temp     Condition")

	for _, day := range days {
		fmt.Printf(
			"%s     %.0fC     %.0f     %.0fC     %s\n",
			day.Date,
			day.Day.MinTempC,
			day.Day.MaxTempC,
			day.Day.AverageTempC,
			day.Day.Condition.Text)

	}
	fmt.Print("\n\n")

}
