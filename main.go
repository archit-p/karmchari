package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// job related information
// jState : to denote current state of job
// jType : to denote the type of job
type jobMeta struct {
	jState int
	jType int
}

// map to store the job states
var jobDict map[string]jobMeta

// a simple way to simulate job table
// tracks currently running jobs
// kills marked jobs
func trackJobs() {
	var prevTime, curTime time.Time

	prevTime = time.Now()
	for {
		curTime = time.Now()
		if (curTime.Sub(prevTime).Seconds() > 1) {
			log.Println("Current jobs")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("|                 ID                 |    State   |    Type    |")
			fmt.Println("----------------------------------------------------------------")

			for it, job := range jobDict {
				fmt.Printf("| %34s | %10s | %10s |\n", it, encodeState(job.jState),
						encodeType(job.jType))
				if job.jState == 2 {
					delete(jobDict, it)
				}
			}

			fmt.Println("----------------------------------------------------------------")
			prevTime = curTime
		}
	}
}

func main() {
	// initialize the jobDict structure
	jobDict = make(map[string]jobMeta)

	// start a routine to track current jobs
	go trackJobs()

	port := "51463"

	log.Printf("Karmchari -- Service started on port %s!", port)

	// create a new mux router and add routes
	r := mux.NewRouter()
	r.HandleFunc("/registerJob", registerJob).Methods("POST")
	r.HandleFunc("/jobState", updateJobState).Methods("POST")
	r.HandleFunc("/jobState", readJobState).Methods("GET")

	// launch the server
	log.Fatalf("Exited: %s", http.ListenAndServe(":" + port, r))

	// exit from the server
	log.Println("Karmchari -- Service stopped!")
}
