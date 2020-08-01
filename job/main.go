package job

// store job metadata
type Job struct {
	// type of job : upload / export / teams
	Type string		`redis:"type"`
	// state of job : start / pause / kill
	State string	`redis:"state"`
}

// returns a new job
func New(Type string) *Job {
	// create a new job and start it
	j := Job{ Type : Type, State : "start" }

	return &j
}

// verify the state string
func VerifyState(State string) bool {
	// valid states
	states := []string{ "start", "pause", "kill" }

	// check whether command is present, and return an encoding
	for _, c := range states {
		if State == c {
			return true
		}
	}

	return false
}

// verify the type string
func VerifyType(Type string) bool {
	// valid types
	types := []string{ "upload", "export", "teams" }

	// check whether command is present, and return an encoding
	for _, t := range types {
		if Type == t {
			return true
		}
	}

	return false
}
