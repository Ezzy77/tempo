/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/Ezzy77/tempo/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get daily weather forecast",
	Long: `This command allows to get daily  
	weather forecast based on the location giving 
	and number of days as an argument.
	Only up to 14 days

	To pass in a location, use the -l or --location
	flag followed by the location name.

	Examples:
	./tempo forecast -l london -d 14
	./tempo forecast --location rome --days 7
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//apiKey := os.Getenv("API_KEY")

		apiKey := viper.GetString("key")
		if apiKey == "" {
			fmt.Println("API key not set. Please run 'tempo setapikey' to set your API key.")

			return nil
		}
		locationName, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Printf("error retrieving location: %s\n",
				err.Error())
			return err
		}
		numberOfDays, err := cmd.Flags().GetInt64("days")
		if err != nil {
			fmt.Printf("error retrieving number of days: %s\n",
				err.Error())
			return err
		}

		numberOfDaysStringVal := strconv.FormatInt(numberOfDays, 10)

		service.GetForecastDays(locationName, apiKey, numberOfDaysStringVal)
		return nil
	},
}

func init() {
	forecastCmd.Flags().StringP("location", "l", "", "location name")
	forecastCmd.Flags().Int64P("days", "d", 0, "number of days")

	forecastCmd.MarkFlagRequired("location")
	forecastCmd.MarkFlagRequired("days")
	rootCmd.AddCommand(forecastCmd)

}
