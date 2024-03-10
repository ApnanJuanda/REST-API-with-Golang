package app

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"katalisStack.com/practice-golang-restful-api/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {

	// ENV Configuration
	err := godotenv.Load("config/.env")
	helper.PanicIfError(err)

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYUSER"),
		os.Getenv("MYPASSWORD"),
		os.Getenv("MYHOST"),
		os.Getenv("MYPORT"),
		os.Getenv("MYDATABASE"))

	db, err := sql.Open("mysql", mysqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
