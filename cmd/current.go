/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Ezzy77/tempo/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current weather information",
	Long: `This command allows to get current 
	weather forecast based on the location giving 
	as argument.

	To pass in a location, use the -l or --location
	flag followed by the location name.

	Examples:
	./tempo current -f london
	./tempo current --location rome
	
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

		service.GetCurrentWeather(locationName, apiKey)
		return nil
	},
}

func init() {
	currentCmd.Flags().StringP("location", "l", "", "location name")
	currentCmd.MarkFlagRequired("location")
	rootCmd.AddCommand(currentCmd)

}
