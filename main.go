package main

import (
	"log"
	"os"
)

func main() {

	var port string

	port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(NewServer(port))

}
