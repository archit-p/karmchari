package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/archit-p/karmchari/jobserver"
)

func main() {
	// configuration parameters for the server
	
	var port, host string
	flag.StringVar(&port, "port", "51463", "port to start app on ex. 51463")
	flag.StringVar(&host, "shost", "redis:6379", "redis host ex. localhost:6379")

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
