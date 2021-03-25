


// Followed: https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://www.sohamkamani.com/golang/parsing-json/

package db

import (
	// "log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" 
	"reflect"
)

const (
    db_port = "3306"
	db = "mydb"
	db_user = "root"
	db_host = "127.0.0.1"
	db_password = "admin"
)

func GetConnection() *sql.DB {
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db))
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	
	return conn
}

func GetType() {
	conn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		reflect.TypeOf(conn)
	}
}