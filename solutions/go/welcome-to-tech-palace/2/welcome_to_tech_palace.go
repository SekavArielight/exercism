package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
	customer = strings.ToUpper(customer)
	return ("Welcome to the Tech Palace," + " " + customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
	return (strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
	oldMsg = strings.Trim(oldMsg, "*")
	return strings.TrimSpace(oldMsg)
}
