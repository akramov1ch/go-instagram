package signin

import (
	"fmt"

	er "github.com/akramov1ch/go-instagram/errors"
	db "github.com/akramov1ch/go-instagram/internal/dbconnect"
	fn "github.com/akramov1ch/go-instagram/internal/functions"
	pr "github.com/akramov1ch/go-instagram/internal/profile"
	st "github.com/akramov1ch/go-instagram/storage/structs"
)

func SignIn() {
	var user st.User
	fmt.Println("Please enter your username")
	fmt.Scanln(&user.Username)
	fmt.Println("Please enter your password")
	fmt.Scanln(&user.Password)
	if fn.CountUsername(user.Username) == 1 {
		if fn.CheckPassword(user.Password, user.Username) == true {
			db, err := db.DbConnect()
			if err != nil {
				panic(err)
			}
			defer db.Close()
			err = db.QueryRow("SELECT user_id, firstname, lastname, email, phone_number, followers, following, followers_list, following_list FROM users WHERE username=$1", user.Username).Scan(&user.User_id, &user.Firstname, &user.Lastname, &user.Email, &user.Phone_number, &user.Followers, &user.Following, &user.Followers_list, &user.Following_list)
			if err != nil {
				panic(err)
			} 
			pr.Profile(user)
		} else {
			fmt.Println(er.ErrorPassword())
			SignIn()
		}
	}
}
