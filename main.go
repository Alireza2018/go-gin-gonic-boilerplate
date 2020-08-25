package main

import (
	"apiservice/router"
	"log"
)

func main() {
	log.Printf("Server started")

	router := router.NewRouter()

	log.Fatal(router.Run(":8080"))
}
