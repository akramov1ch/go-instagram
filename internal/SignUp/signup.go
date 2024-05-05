package signup

import (
	"fmt"

	db "github.com/akramov1ch/go-instagram/internal/dbconnect"
	fn "github.com/akramov1ch/go-instagram/internal/functions"
	st "github.com/akramov1ch/go-instagram/storage/structs"
	er "github.com/akramov1ch/go-instagram/errors"
)

func SignUp() () {
	var newuser st.User
	fmt.Println("Please enter your username")
	fmt.Scanln(&newuser.Username)
	if fn.CountUsername(newuser.Username) == 0 {
		if fn.CheckUsername(newuser.Username) == true{
			fmt.Println("Please enter your email")
			fmt.Scanln(&newuser.Email)
			if fn.CountEmail(newuser.Email) == true && fn.CheckEmail(newuser.Email) == true {
				fmt.Println("Please enter your password")
				fmt.Scanln(&newuser.Password)
				if len(newuser.Password) >= 8 {
					newuser.Password = fn.HashPassword(newuser.Password)
					db, err := db.DbConnect()
					fmt.Println("Please enter your first name")
					fmt.Scanln(&newuser.Firstname)
					fmt.Println("Please enter your last name")
					fmt.Scanln(&newuser.Lastname)
					fmt.Println("Please enter your phone number")
					fmt.Scanln(&newuser.Phone_number)
					fmt.Println("Please wait while we create your account")
					if err!= nil {
						panic(err)
					}
					defer db.Close()
					err = db.QueryRow("INSERT INTO users (username, email, password, firstname, lastname, phone_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id, username, firstname, lastname, email, password, phone_number", newuser.Username, newuser.Email, newuser.Password, newuser.Firstname, newuser.Lastname, newuser.Phone_number).Scan(&newuser.User_id, &newuser.Username, &newuser.Firstname, &newuser.Lastname, &newuser.Email, &newuser.Password, &newuser.Phone_number)
					if err!= nil {
						panic(err)
					}
				} else {
					fmt.Println(er.ErrorPassword())
					SignUp()
				}
			} else {
				fmt.Println(er.HaveEmail())
				SignUp()
			}
		} else {
			fmt.Println(er.ErrorUsername())
			SignUp()
		}
	} else {
		fmt.Println(er.HaveUsername())
		SignUp()
	}
}
