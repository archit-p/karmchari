package repo

type Job struct {
	Type string		`redis:"type"`
	State string	`redis:"state"`
}

func NewJob(JobType string) *Job {
	j := Job{ Type : JobType, State : "start" }

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
