package service

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
)


var db *sql.DB = nil

func GetDB() (*sql.DB, error){

	if db == nil{
		host := "35.199.58.216"
		uname := "root"
		pwd := "BBc-7rxp"
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/hophacks", uname, pwd, host)

		log.Println(connectionString)

		dbConn, err := sql.Open("mysql", connectionString)
		if err != nil{
			return nil, err
		}

		pingErr := dbConn.Ping()
		if pingErr != nil{
			return nil, pingErr
		}


		//db is the package level db
		db = dbConn
	}

	return db, nil
}