package utils

import (
	"errors"
)

type userPass struct {
	username string
	password string
}

func TryLogin(username string, password string) (string, error) {
	// user1 instance
	user1 := userPass{
		username: "opet",
		password: "opet123",
	}

	// user2 instance
	user2 := userPass{
		username: "piyan",
		password: "piyan123",
	}

	// user3 instance
	user3 := userPass{
		username: "ce",
		password: "ce123",
	}

	var database []userPass = []userPass{
		user1,
		user2,
		user3,
	}

	// Loop through entire database
	for _, user := range database {
		if username == user.username {
			if password == user.password {
				successMsg := "Selamat datang, " + username + "!"
				return successMsg, nil
			}
			return "", errors.New("User atau pass salah (actual error : pass tidak match)")
		}
	}

	return "", errors.New("User atau pass salah (actual error : user tidak ditemukan)")
}
