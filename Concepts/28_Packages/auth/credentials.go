package auth

import "fmt"

// must use Capital letter for naming Function, if use small letter, then cann't use outside the folder(private)

func LoginWithCredentials(username string, password string) {
	fmt.Println("login with user", username, password)
}
