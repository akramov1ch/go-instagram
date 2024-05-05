package dbconnect

import(
	"database/sql"
	db "github.com/akramov1ch/go-instagram/storage/database"
	_ "github.com/lib/pq"
)


func DbConnect()(*sql.DB, error){
	dbinfo := db.DbInfo()
	db, err := sql.Open("postgres", dbinfo)
	if err!= nil{
		return nil, err
	}
	return db, nil
}