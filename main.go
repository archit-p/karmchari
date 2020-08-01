package main

import (
	"log"
	"net/http"

	"github.com/archit-p/karmchari/jobserver"
)

func main() {
	// configuration parameters for the server
	port := "51463"
	host := "redis:6379"

	// initialize a new server
	js, err := jobserver.New(port, host)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Karmchari -- Service started on port %s!", port)

	// launch the server
	log.Fatalf("Exited: %s", http.ListenAndServe(":" + js.ServerPort, js.Router))

	// exit from the server
	log.Println("Karmchari -- Service stopped!")
}
