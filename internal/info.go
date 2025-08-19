package personalInfo

import "fmt"

func PrintInfo() {
	type template struct {
		nama    string
		foto    string
		email   string
		umur    uint8
		noTelp  string
		isNikah bool
	}

	user1 := template{
		nama:    "Opet",
		foto:    "/opet.jpg",
		email:   "opet@gmail.com",
		umur:    20,
		noTelp:  "0851551357221",
		isNikah: false,
	}

	fmt.Println(user1)

}
