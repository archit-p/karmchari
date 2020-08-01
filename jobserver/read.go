package jobserver

import (
	"fmt"
	"net/http"
)

// allows for reading the current state of job
func (js *JobServer) ReadJobHandler(w http.ResponseWriter, r *http.Request) {
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
	curJob, err := js.Store.Read(jobId)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Job not available")
		return
	}

	// return success along with state
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s -- %s", jobId, curJob.State)
}
