package dbconnect

import(
	"database/sql"
	_ "github.com/lib/pq"
)

func DbConnect(dbinfo string) (*sql.DB, error){
	db, err := sql.Open("postgres", dbinfo)
	if err!= nil{
		return nil, err
	}
	return db, nil
}