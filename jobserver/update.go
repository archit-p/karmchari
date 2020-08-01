package jobserver

import (
	"fmt"
	"net/http"

	"github.com/archit-p/karmchari/job"
)

// allows for stopping / resuming / killing a job
func (js *JobServer) UpdateJobHandler(w http.ResponseWriter, r *http.Request) {
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

	// extract new state
	stateCommand := r.FormValue("command")
	if stateCommand == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Command is required")
		return
	}

	// check whether the state is valid
	if !job.VerifyState(stateCommand) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Command is invalid")
	}

	err := js.Store.UpdateState(jobId, stateCommand)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Failed to update")
		return
	}

	// return success
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s -- %s", jobId, stateCommand)
}
