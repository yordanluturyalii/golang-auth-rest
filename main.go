package main

import (
	"log"
	"yordanluturyali/golang-auth-rest/app"
)

func main() {
	fiberApp := app.NewRouter();

	if err := fiberApp.Listen("127.0.0.1:3000"); err != nil {
		log.Fatal("Error: ", err)
		panic(err)
	}
}
