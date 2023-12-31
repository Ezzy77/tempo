/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/Ezzy77/tempo/cmd"
	"github.com/Ezzy77/tempo/config"
	"github.com/Ezzy77/tempo/internal"
	"github.com/spf13/viper"
)

func main() {
	//godotenv.Load()

	// create file to store weather api key locally
	internal.CreateViperConfigFile()

	// init viper configuration
	viper.SetConfigName("tempoCLI_Config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME")
	//viper.AutomaticEnv()

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&config.Config); err != nil {
		fmt.Println("Error unmarshaling config:", err)
	}
	cmd.Execute()
}
