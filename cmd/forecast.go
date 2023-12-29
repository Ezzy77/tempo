/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast",
	Long: `This command allows to to get 10 days  
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
		locationName, err := cmd.Flags().GetString("location")
		if err != nil {
			fmt.Printf("error retrieving location: %s\n",
				err.Error())
			return err
		}

		if locationName == "" {
			fmt.Println("showing weather forecast at your current location")
			return nil
		}
		fmt.Println("showing weather forecast for " + locationName)
		return nil
	},
}

func init() {
	forecastCmd.Flags().StringP("location", "l", "", "location name")
	rootCmd.AddCommand(forecastCmd)

}
