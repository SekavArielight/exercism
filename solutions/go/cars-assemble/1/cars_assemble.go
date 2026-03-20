package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	return  float64(productionRate) * successRate/100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	// return productionRate/60 * int(successRate/100)
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	var groups uint = uint(carsCount/10)
	var remainder int = carsCount % 10
	return uint(groups * 95000) + uint(remainder * 10000)
}
