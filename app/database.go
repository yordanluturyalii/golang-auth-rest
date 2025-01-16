package app

import (
	"database/sql"
	"log"
	"yordanluturyali/golang-auth-rest/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(cnf *config.Config) *sql.DB {
	dbHost := cnf.Env.GetString("DB_HOST")
	dbPort := cnf.Env.GetString("DB_PORT")
	dbUser := cnf.Env.GetString("DB_USERNAME")
	dbPass := cnf.Env.GetString("DB_PASS")
	dbName := cnf.Env.GetString("DB_NAME")

	dns := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	
	con, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("error while connecting to database", err)
	}

	err = con.Ping()
	if err != nil {
		log.Fatal("error while ping to database", err)
	}

	return con
}