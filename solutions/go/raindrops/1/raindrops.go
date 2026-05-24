package raindrops

import "strconv"

func Convert(number int) string {
	// panic("Please implement the Convert function")
	// if number % 3 == 0 {
	// 	return "Pling"
	// }
	// if number % 5 == 0 {
	// 	return "Plang"
	// }
	// if number % 7 == 0 {
	// 	return "Plong"
	// }

	switch {
		case number % 3 == 0 && number % 5 == 0 && number % 7 == 0:
			return "PlingPlangPlong"
		case number % 3 == 0 && number % 5 == 0:
			return "PlingPlang"
		case number % 3 == 0 && number % 7 == 0:
			return "PlingPlong"
		case number % 5 == 0 && number % 7 == 0:
			return "PlangPlong"
		case number % 3 == 0:
			return "Pling"
		case number % 5 == 0:
			return "Plang"
		case number % 7 == 0:
			return "Plong"	
	}
	return strconv.Itoa(number)
}
