package jobserver

import (
	"fmt"
	"net/http"

	"github.com/archit-p/karmchari/job"
)

// allows for adding a new job
func (js *JobServer) RegisterJobHandler(w http.ResponseWriter, r *http.Request) {
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
	if !job.VerifyType(jType) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Type is invalid")
		return
	}

	id, err := js.Store.Write(jType, "start")
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Failed")
		fmt.Fprint(w, "Failed to write job")
		return
	}

	// return success along with pid
	w.WriteHeader(200)
	fmt.Fprintf(w, "Created! id = %s", *id)
}
