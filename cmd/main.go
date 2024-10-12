package main

import "teste/internal/adapters/http"

func main() {
	e := http.NewWebService()
	e.Logger.Fatal(e.Start(":8080"))
}
