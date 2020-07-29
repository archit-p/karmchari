package main

import (
	"fmt"
	"net/http"
)

// allows for adding a new job
func registerJob(w http.ResponseWriter, r *http.Request) {
	// parse the POST fields, for accessing them through FormValue
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal error")
		return
	}

	// extract job type from form fields
	jType := r.FormValue("type")
	if jType == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Type is required")
		return
	}

	// decode the type string into an integer
	jTypeDec := decodeType(jType)
	if jTypeDec == -1 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Type is invalid")
	}

	// generate hash for use as job ID
	pid := getHashString()

	// create an entry
	jobDict[pid] = jobMeta{ jState : 0, jType : jTypeDec }

	// return success along with pid
	w.WriteHeader(200)
	fmt.Fprintf(w, "Created! id = %s", pid)
}

// allows for reading the current state of job
func readJobState(w http.ResponseWriter, r *http.Request) {
	// parse the GET fields
	q := r.URL.Query()

	jobIds := q["id"]
	if len(jobIds) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "ID is required")
		return
	}
	jobId := jobIds[0]

	// disallow requests with illegal job ID
	curJob, ok := jobDict[jobId]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Job not available")
		return
	}

	// return success along with state
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s -- %s", jobId, encodeState(curJob.jState))
}

// allows for stopping / resuming / killing a job
func updateJobState(w http.ResponseWriter, r *http.Request) {
	// parse the POST fields, for accessing them through FormValue
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal error")
		return
	}

	// extract job id from form fields
	jobId := r.FormValue("id")
	if jobId == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "ID is required")
		return
	}

	// disallow requests with illegal job ID
	curJob, ok := jobDict[jobId]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Job not available")
		return
	}

	// extract new state
	stateCommand := r.FormValue("command")
	if stateCommand == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Command is required")
		return
	}

	// check whether the state is valid
	decState := decodeState(stateCommand)
	if decState == -1 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Command is invalid")
	}

	// update the job state
	curJob.jState = decState
	jobDict[jobId] = curJob

	// return success
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s -- %s", jobId, stateCommand)
}
