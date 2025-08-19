package main

import (
	"fmt"

	"github.com/radifan9/minitask-w8-d1/internal/utils"
)

func main() {
	fmt.Println("Minitask W8 D1")

	// No 1
	// utils.PrintInfo()

	// No 2
	success, err := utils.TryLogin("opet", "opet123")
	// Check wether there's an error or not. If there's an error, err should not be nil
	if err == nil {
		// Logged in sucess
		fmt.Println(success)
	} else {
		// Logged in failed
		fmt.Println(err.Error())
	}
}
