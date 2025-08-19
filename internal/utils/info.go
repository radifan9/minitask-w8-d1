package utils

import "fmt"

// Define a struct for individual school
type school struct {
	nama    string
	jurusan string
}

// template represents user data including personal information and educational background
type template struct {
	nama       string
	foto       string
	email      string
	umur       uint8
	noTelp     string
	isNikah    bool
	listSchool []school
}

// PrintInfo demonstrates the usage of the user data structure
func PrintInfo() {

	// Create some school instances
	schoolSMK := school{
		nama:    "SMK 1 Surabaya",
		jurusan: "TKJ",
	}

	schoolKuliah := school{
		nama:    "Universitas Indonesia",
		jurusan: "Computer Science",
	}

	// Create template instance and populate it's listSchool slice
	user1 := template{
		nama:       "Opet",
		foto:       "/opet.jpg",
		email:      "opet@gmail.com",
		umur:       20,
		noTelp:     "0851551357221",
		isNikah:    false,
		listSchool: []school{schoolSMK, schoolKuliah},
	}

	// Print user information
	fmt.Printf("--- --- No 1 --- ---\n")
	fmt.Printf("User Information:\n")
	fmt.Printf("Name: %s\n", user1.nama)
	fmt.Printf("Age: %d\n", user1.umur)
	fmt.Printf("Email: %s\n", user1.email)
	fmt.Printf("Phone: %s\n", user1.noTelp)
	fmt.Printf("Married: %t\n", user1.isNikah)
	fmt.Printf("Photo: %s\n", user1.foto)
	fmt.Printf("Education:\n")
	for i, school := range user1.listSchool {
		fmt.Printf("  %d. %s - %s\n", i+1, school.nama, school.jurusan)
	}

}
