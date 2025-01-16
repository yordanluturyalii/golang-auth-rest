package main

import (
	"log"
	"yordanluturyali/golang-auth-rest/app"
	"yordanluturyali/golang-auth-rest/config"
)

func main() {
	cnf := config.NewConfig()
	_ = app.NewDatabase(cnf)
	fiberApp := app.NewRouter();

	if err := fiberApp.Listen("127.0.0.1:3000"); err != nil {
		log.Fatal("Error: ", err)
		panic(err)
	}
}
