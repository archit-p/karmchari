package main

// decode the state string into an integer
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

// decode the state string into an integer
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
