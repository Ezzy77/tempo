/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Ezzy77/tempo/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Print the location information
	//fmt.Printf("City: %s\nRegion: %s\nCountry: %s\nLocation: %s\n", location.City, location.Region, location.Country, location.Loc)

	cmd.Execute()
}
