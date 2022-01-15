package db

import (
	"database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func Connect() (*sql.DB,error) {
	dbConn, err := sql.Open("mysql", "balahush:Stud_21g@/restaurant_db")
	return dbConn, err
}