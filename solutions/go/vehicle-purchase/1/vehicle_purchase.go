package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 +  " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
	if age < 3 {
		return (80 * originalPrice)/100
	} else if age >= 10 {
		return (50 * originalPrice)/100
	// } else if age >= 3 && age < 10 {
	} else {
		return (70 * originalPrice)/100
	}
	// return originalPrice
}
