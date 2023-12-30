/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Ezzy77/tempo/internal"
	"github.com/Ezzy77/tempo/service"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast",
	Long: `This command allows to to get 14 days  
	weather forecast based on the location giving 
	as an argument. If no argument is given the default
	forcast will be your current location.

	To pass in a location, use the -l or --location
	flag followed by the location name.

	Examples:
	./tempo forecast -f london
	./tempo forecast --location rome
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("API_KEY")

		locationName, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Printf("error retrieving location: %s\n",
				err.Error())
			return err
		}

		if locationName == "" {
			// Get the public IP address of the machine
			ip, err := internal.GetPublicIP()
			if err != nil {
				fmt.Println("Error getting public IP:", err)
				return err
			}

			// Use an IP geolocation API to get location information
			location, err := internal.GetLocationInfo(ip)
			if err != nil {
				fmt.Println("Error getting location information:", err)
				return err
			}
			service.GetForecast(location.City, apiKey)

			return nil
		}

		service.GetForecast(locationName, apiKey)
		return nil
	},
}

func init() {
	forecastCmd.Flags().StringP("location", "l", "", "location name")
	rootCmd.AddCommand(forecastCmd)

}
