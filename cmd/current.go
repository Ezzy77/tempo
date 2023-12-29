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
	weather forcast based on the location giving 
	as argument. If no argument given the default
	forcast will be your current location.

	To pass in a location, use the -l or --location
	flag followed by the location name.

	Examples:
	./tempo current -f london
	./tempo current --location rome
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("current called")
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

}
