// Package weather provides weather forecast functionality.
package weather

var (
	CurrentCondition string
	CurrentLocation  string
)

// Forecast returns the current weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	// CurrentLocation is updated with the current location.
	// CurrentCondition is updated with the current condition.
	CurrentLocation, CurrentCondition = city, condition
	
	// Return the current location and condition.
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
