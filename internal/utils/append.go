package utils

import "fmt"

func AppendSlice() {
	sliceOfInt := []int{50, 75, 66, 20, 32, 90}
	insertIndex := 3

	// Split the slice
	firstPart := sliceOfInt[:insertIndex]
	secondPart := sliceOfInt[insertIndex:]

	sliceOfInt = append(firstPart, append([]int{88}, secondPart...)...)

	fmt.Println(sliceOfInt)
}
