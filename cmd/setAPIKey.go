/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Ezzy77/tempo/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setAPIKeyCmd represents the setAPIKey command
var setAPIKeyCmd = &cobra.Command{
	Use:   "setapikey",
	Short: "Set API key for weather forecast",
	Long: `This command allows user to set their API Key
	giving by www.weatherapi.com in order to be able to
	have access to weather information.

	To pass in an API key, use the -k or --key
	flag followed by the API key value.

	Examples:
	./tempo setapikey -k <your api key string>
	./tempo setapikey --key <your api key string>
.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Printf("error retrieving api key: %s\n",
				err.Error())
			return err
		}
		// Update configuration with the new API Key
		config.Config.Key = apiKey

		// Save the updated config to the file
		errViper := viper.WriteConfig()
		if err != nil {
			fmt.Println(errViper)
		}

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error writing config file: %s\n",
				err.Error())
			return err
		}

		fmt.Println("API key set successfully.")
		return nil
	},
}

func init() {
	setAPIKeyCmd.Flags().StringP("key", "k", "", "Weather API KEY")
	viper.BindPFlag("key", setAPIKeyCmd.Flags().Lookup("key"))

	rootCmd.AddCommand(setAPIKeyCmd)

}
