/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
	./tempo forecast -l london -d 
	./tempo forecast --location rome --days 7
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("API_KEY")

		locationName, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Printf("error retrieving location: %s\n",
				err.Error())
			return err
		}
		numberOfDays, err := cmd.Flags().GetString("days")
		if err != nil {
			fmt.Printf("error retrieving number of days: %s\n",
				err.Error())
			return err
		}

		service.GetForecastDays(locationName, apiKey, numberOfDays)
		return nil
	},
}

func init() {
	forecastCmd.Flags().StringP("location", "l", "", "location name")
	forecastCmd.Flags().StringP("days", "d", "", "number of days")

	forecastCmd.MarkFlagRequired("location")
	forecastCmd.MarkFlagRequired("days")
	rootCmd.AddCommand(forecastCmd)

}
