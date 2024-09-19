package weather

import (
	"fmt"
	"strings"
)

const KelvinSubtract = 273.15

// Coords represents a physical place on planet earth.
type Coords struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Weather represents a type of weather at a certain location
type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Data was originally generated via GoLand's json -> type definition generator.
// It has been modified for cleanliness.
type Data struct {
	Coords  Coords    `json:"coord"`
	Weather []Weather `json:"weather"`
	Base    string    `json:"base"`
	Main    struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone   int    `json:"timezone"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
	unitSymbol string
}

// ToCelsius converts all the temperature ratings returned from the API from kelvin to Celsius.
func (d Data) ToCelsius() Data {
	d.Main.Temp = d.Main.Temp - KelvinSubtract
	d.Main.TempMin = d.Main.TempMin - KelvinSubtract
	d.Main.TempMax = d.Main.TempMax - KelvinSubtract
	d.Main.FeelsLike = d.Main.FeelsLike - KelvinSubtract

	d.unitSymbol = "°C"

	return d
}

func toFahrenheit(k float64) float64 {
	return k*9/5 - 459.67
}

// ToFreedomUnits converts all the temperature ratings from the API from kelvin to freedom units (Fahrenheit)
func (d Data) ToFreedomUnits() Data {
	d.Main.Temp = toFahrenheit(d.Main.Temp)
	d.Main.TempMin = toFahrenheit(d.Main.TempMin)
	d.Main.TempMax = toFahrenheit(d.Main.TempMax)
	d.Main.FeelsLike = toFahrenheit(d.Main.FeelsLike)

	d.unitSymbol = "°F"

	return d
}

func (d Data) String() string {
	if d.unitSymbol == "" {
		d.unitSymbol = "K"
	}

	return strings.ReplaceAll(fmt.Sprintf(`== Weather In: %s ==
Weather: %s
Temperature: %.2f[unit] (min: %.2f[unit], max: %.2f[unit])
Feels Like: %.2f[unit]
Humdity: %d%%`, d.Name, d.Weather[0].Main, d.Main.Temp, d.Main.TempMin, d.Main.TempMax, d.Main.FeelsLike, d.Main.Humidity), "[unit]", d.unitSymbol)
}

type Provider interface {
	Get(location string) (Data, error)
}
