package database

import "fmt"

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "vakhaboff"
	dbname = "instagram"
	sslmode = "disable"
)

func DbInfo() string{
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	return dbinfo
}