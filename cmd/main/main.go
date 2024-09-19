package main

import (
	"flag"
	"fmt"
	"github.com/liondadev/csp-weather-app/weather"
	"os"
	"strings"
)

func printUsageAndExit() {
	fmt.Println("Usage: weather-cli [--unit kcf] (location)")
	os.Exit(1)
}

var apiKey = os.Getenv("OPENWEATHERMAP_API_KEY")
var provider weather.Provider

func init() {
	if apiKey == "" {
		panic("Environment variable OPENWEATHERMAP_API_KEY is not set, please set it to your open weather map api key.")
	}

	provider = weather.NewOpenWeatherMapProvider(apiKey)
}

func main() {
	// os.Args contains the program's full path in the first argument
	// we don't need that, so we discard it completely
	args := os.Args[1:]

	if len(args) < 1 {
		printUsageAndExit()
		return
	}

	// Allow people to put a flag in to change the unit
	var unit string
	flag.StringVar(&unit, "unit", "c", "the unit to return the temperature in (c, k, or f)")

	flag.Parse()

	if strings.HasPrefix(args[0], "--") {
		// If the flag is formatted like --<name>=<val>, we don't want to skip over two
		if strings.Contains(args[0], "=") {
			args = args[1:]
		} else {
			args = args[2:]
		}
	}

	// Get the weather data
	data, err := provider.Get(strings.Join(args, " "))
	if err != nil {
		panic(err)
	}

	// Convert the weather data to the requested unit
	switch strings.ToLower(unit) {
	case "c":
		data = data.ToCelsius()
		break
	case "freedom":
		fallthrough
	case "f":
		data = data.ToFreedomUnits()
		break
	default:
		fmt.Printf("Warning! Unknown unit '%s'. Defaulting to kelvin.\n", unit)
	}

	// Print the output
	fmt.Println(data.String())
}
