package utils

import "fmt"

func Find() uint8 {
	fmt.Printf("\n\n--- --- No 3 --- ---\n")

	arrayPositiveNum := [5]uint8{5, 10, 27, 20, 25}

	largest := uint8(0)
	for _, value := range arrayPositiveNum {
		// fmt.Println(value)
		if value > uint8(largest) {
			largest = value
		}
	}
	return largest
}
