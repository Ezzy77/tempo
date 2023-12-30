package models

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
			Date string `json:"date"`
			Day  struct {
				MaxTempC          float64 `json:"maxtemp_c"`
				MinTempC          float64 `json:"mintemp_c"`
				AverageTempC      float64 `json:"avgtemp_c"`
				DailyChanceOfRain string  `json:"daily_chance_of_rain"`
				Condition         struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
