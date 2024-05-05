package main

import (
	"fmt"
	"strings"

	si "github.com/akramov1ch/go-instagram/internal/SignIn"
	su "github.com/akramov1ch/go-instagram/internal/SignUp"
)

func main() {
    topBorder := strings.Repeat("=", 40)
    bottomBorder := strings.Repeat("=", 40)
    separator := strings.Repeat("-", 40)
    fmt.Println(topBorder)
    fmt.Println("| Welcome to Go-Instagram |")
    fmt.Println(separator)
    fmt.Println("| Please select an option |")
    fmt.Println(separator)
    fmt.Println("| 1. Sign up              |")
    fmt.Println("| 2. Sign in              |")
    fmt.Println("| 3. Exit                 |")
    fmt.Println(bottomBorder)
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		su.SignUp()
	case 2:
		si.SignIn()
	case 3:
		fmt.Println("Thank you for using Go-Instagram")
	default:
		fmt.Println("Invalid option")
		main()
	}
}
