/*
Package data - Handles functions related to data source access e.g. cache, databases
*/
package data

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Points to the writeable blogs database.
var BlogDB *sql.DB

// IST : set dafault time zone
var IST *time.Location

// InitDB initialises the database pools with
func InitDB(host, port, user, password, readhost, readport, readuser, readpassword string) (BlogDB *sql.DB, err error) {
	BlogDB, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/blogs")

	if err != nil {
		return nil, errors.New("DataBase error :" + err.Error())
	}
	if err = BlogDB.Ping(); err != nil {
		return nil, errors.New("DataBase error :" + err.Error())
	}

	IST, _ = time.LoadLocation("Asia/Kolkata")

	return BlogDB, nil
}
