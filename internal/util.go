package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

// GeoLocation represents the structure of the geolocation response.
type GeoLocation struct {
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Loc     string `json:"loc"`
}

// getPublicIP gets the public IP address of the machine.
func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api64.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	ip := data["ip"]
	return ip, nil
}

// getLocationInfo gets location information based on the provided IP address.
func GetLocationInfo(ip string) (GeoLocation, error) {
	resp, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		return GeoLocation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GeoLocation{}, err
	}

	var location GeoLocation
	if err := json.Unmarshal(body, &location); err != nil {
		return GeoLocation{}, err
	}

	return location, nil
}

// GETTING CURRENT LOCATION IF -LOCATION IS NOT GIVEN BY THE USER
// THIS FEATURE IS REMOVED FROM THE FORECAST COMMAND  DUE TO BEING SLOW
// if locationName == "" {
// 	// Get the public IP address of the machine
// 	ip, err := internal.GetPublicIP()
// 	if err != nil {
// 		fmt.Println("Error getting public IP:", err)
// 		return err
// 	}

// 	// Use an IP geolocation API to get location information
// 	location, err := internal.GetLocationInfo(ip)
// 	if err != nil {
// 		fmt.Println("Error getting location information:", err)
// 		return err
// 	}
// 	service.GetForecast(location.City, apiKey, numberOfDays)

// 	return nil
// }
