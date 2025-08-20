package utils

import "fmt"

func ConvertTemp(suhuInput int16, skalaInput string, skalaTujuan string) (int16, string, error) {
	fmt.Printf("\n\n--- --- No 4 --- ---\n")

	konstantaSkala := map[string]int16{
		"C": 5,
		"R": 4,
		"K": 5,
		"F": 9,
	}

	freezingPoint := map[string]int16{
		"C": 0,
		"R": 0,
		"F": 32,
		"K": 273,
	}

	// if skalaInput == "K" {
	// 	suhuInput = suhuInput - 273
	// }

	// if skalaInput == "F" {
	// 	suhuInput = suhuInput - 32
	// }

	suhuTujuan := (konstantaSkala[skalaTujuan]/konstantaSkala[skalaInput])*(suhuInput-freezingPoint[skalaInput]) + freezingPoint[skalaTujuan]

	// if skalaTujuan == "K" {
	// 	suhuTujuan = suhuTujuan + 273
	// }

	// if skalaTujuan == "F" {
	// 	suhuTujuan = suhuTujuan + 32
	// }

	return suhuTujuan, skalaTujuan, nil
}
