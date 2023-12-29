/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current weather information",
	Long: `This command allows to to get current 
	weather forecast based on the location giving 
	as argument. If no argument given the default
	forcast will be your current location.

	To pass in a location, use the -l or --location
	flag followed by the location name.

	Examples:
	./tempo current -f london
	./tempo current --location rome
	
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
		fmt.Println("showing current weather info for " + locationName)
		return nil
	},
}

func init() {
	currentCmd.Flags().StringP("location", "l", "", "location name")
	rootCmd.AddCommand(currentCmd)

}
