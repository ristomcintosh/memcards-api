package main

import (
	"fmt"
	"log"
	"net/http"
)

type application struct {}

func main() {
	app := application{}

	handlers := app.routes()

	fmt.Println("listening on port 5757")
	log.Fatal(http.ListenAndServe(":5757", handlers))
}