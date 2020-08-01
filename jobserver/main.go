package jobserver

import (
	"github.com/archit-p/karmchari/jobstore"

	"github.com/gorilla/mux"
)

type JobServer struct {
	// port to serve on
	ServerPort string
	// host to interact with redis on
	StoreHost string
	// jobstore for the interfacing with redis
	Store *jobstore.JobStore
	// router for setting the API routes
	Router *mux.Router
}

func New(port string, host string) (*JobServer, error) {
	js := JobServer { ServerPort : port, StoreHost : host }

	// initialize a new jobstore struct
	j, err := jobstore.New(js.StoreHost)
	if err != nil {
		return nil, err
	}

	// pass it into the jobserver struct
	js.Store = j
	// create a new mux router
	r := mux.NewRouter()

	// set routes
	r.HandleFunc("/registerJob", js.RegisterJobHandler).Methods("POST")
	r.HandleFunc("/jobState", js.UpdateJobHandler).Methods("POST")
	r.HandleFunc("/jobState", js.ReadJobHandler).Methods("GET")

	// pass it into the jobserver struct
	js.Router = r

	return &js, nil
}
