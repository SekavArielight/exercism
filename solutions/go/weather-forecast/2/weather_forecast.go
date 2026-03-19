// Package weather provides weather forecast functionality.
package weather

var (
	// CurrentCondition is updated with the current condition.
	CurrentCondition string
	// CurrentLocation is updated with the current location.
	CurrentLocation  string
)

// Forecast returns the current weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
