package main

import (
	"fmt"

	"github.com/radifan9/minitask-w8-d1/internal/utils"
)

func main() {
	fmt.Println("Minitask W8 D1")

	// --- --- No 1 --- ---
	utils.PrintInfo()

	// --- --- No 2 --- ---
	success, err := utils.TryLogin("opet", "opet123")
	// Check wether there's an error or not. If there's an error, err should not be nil
	if err == nil {
		// Logged in sucess
		fmt.Println(success)
	} else {
		// Logged in failed
		fmt.Println(err.Error())
	}

	// --- --- No 3 --- ---
	largestNum := utils.Find()
	fmt.Println("Largest Number : ", largestNum)

	// --- --- No 4 --- ---
	convertedTemp, convertedScale, error := utils.ConvertTemp(100.0, "C", "F")
	fmt.Println(convertedTemp)
	fmt.Println(convertedScale)
	fmt.Println(error)

	// Should be 373.15
	fmt.Printf("Suhu hasil konversi : %f %s\n", float64(convertedTemp), convertedScale)

	// --- --- No 5 --- ---
	utils.AppendSlice()
}
