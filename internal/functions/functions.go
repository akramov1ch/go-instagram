package functions

import (
	"crypto/sha256"
	"fmt"
	"unicode"

	db "github.com/akramov1ch/go-instagram/internal/dbconnect"
)

func CountText(text string, chr string) int {
	cnt := 0
	for _, val := range text{
		if string(val) ==  chr{
			cnt++
		}
	}
	return cnt
}

func CountUsername(username string) int {
	db, err := db.DbConnect()
	if err!= nil {
		panic(err)
	}
	defer db.Close()
	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM users WHERE username=$1", username).Scan(&count)
	return count
}

func CountEmail(email string) bool {
	db, err := db.DbConnect()
	if err!= nil {
		panic(err)
	}
	defer db.Close()
	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM users WHERE email=$1", email).Scan(&count)
	if count == 0 {
		return true
	} else {
		return false
	}
}

func CheckEmail(email string) bool {
	if CountText(email, "@") == 1 && CountText(email, ".") == 1 && CountText(email, " ") == 0 {
		return true
	} else {
		return false
	}
}


func HashPassword(password string) string {
    pass_byte := []byte(password)
    pass_shifr := sha256.Sum256(pass_byte)
    hashHex := fmt.Sprintf("%x", pass_shifr)
    return hashHex
}

func CheckPassword(password string, username string) bool {
	db, err := db.DbConnect()
	if err!= nil {
		panic(err)
	}
	defer db.Close()
	var hashHex string
	_ = db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&hashHex)
	if hashHex == HashPassword(password) {
		return true
	} else {
		return false
	}
}

func CheckUsername(username string) bool {
	if len(username) == 0 {
		return false
	}
	for _, char := range username {
		if !unicode.IsLower(char) && !unicode.IsDigit(char) && char != '.' && char != '_' {
			return false
		}
	}
	return true
}