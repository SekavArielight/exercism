package cards

// cards := []int {2, 6, 9}
// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	// panic("Please implement the FavoriteCards function")
	cards := []int {2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// panic("Please implement the GetItem function")
	for i, v := range slice {
		if i == index {
			return v
		}
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	// panic("Please implement the SetItem function")
	if index < 0 || index >= len(slice) {
		return append(slice, value)
	}
	slice[index] = value
	return slice
	
	
	

	// if index >= 0 || index < len(slice) {
	// 	slice[index] = value
	// 	return slice
	// } else if index > len(slice) || index < 0 {
	// 	return append(slice, value)
	// }
	// return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// panic("Please implement the PrependItems function")
	// newSlice := slice[]
	slice = append(values, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// panic("Please implement the RemoveItem function")
	if index < 0 || index > len(slice) {
		return slice
	}
	firstIndex := slice[:index]
	lastIndex := slice[index+1:]
	return append(firstIndex, lastIndex...)
}
