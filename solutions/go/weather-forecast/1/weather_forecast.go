// Package weather provides s function that lets you forecast weather conditions for different cities.
package weather

var (
    // CurrentCondition represents the actual weather condition.
	CurrentCondition string
    // CurrentLocation represents the Location.
	CurrentLocation  string
)

// Forecast function forecasts the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
