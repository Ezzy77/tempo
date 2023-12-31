# tempo
tempo is a command-line tool built in Go using the Cobra framework for checking weather forecasts. Stay informed about the current weather conditions and upcoming forecasts right from your terminal.

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/your-repo)](https://goreportcard.com/report/github.com/yourusername/your-repo)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# Features
- **Current Weather:** Get real-time information about the current weather conditions in your city or worldwide.
- **Forecast:** Retrieve a detailed weather forecast for up to 14 days.
- **Location-Based:** Search for weather information based on location given as an argument.

## Table of Contents

- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Usage

To use tempo-cli on your machine follow these steps:
1. Make sure you get your weather API Keys from https://www.weatherapi.com/ you will be prompt to provide your API key when using the CLI 
2. Clone the repository into your local machine and make sure you have Go installed
3. Run the application with the commands you want to execute

## Commands

1. To set your Weather API Key run this command if the project is opened in VSCode:
```bash
     go run main.go setapikey -k <YOUR_API_KEY>
```
After setting your API key, the key will be saved in a file in your local machine

2. Getting current weather information of provided city, run:
```bash
     go run main.go current -l <CITY>
```
3. Getting weather forecast up to 14 days, run:
```bash
     go run main.go forecast -l <CITY> -d <NUMBER OF DAY>
```

## Examples
# Setting your weather API key: 

![Alt Text](https://github.com/Ezzy77/tempo/blob/main/images/settingapikey.png)


# Getting current weather information of provided city: 

![Alt Text](https://github.com/Ezzy77/tempo/blob/main/images/current.png)

# Getting weather forecast up to 14 days: 

![Alt Text](https://github.com/Ezzy77/tempo/blob/main/images/forecast.png)
