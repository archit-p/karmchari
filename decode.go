package main

// decode the state string into an integer
func decodeState(command string) int {
	// valid states
	states := []string{ "start", "pause", "kill" }

	// check whether command is present, and return an encoding
	for it, c := range states {
		if command == c {
			return it
		}
	}

	return -1
}

// decode the state string into an integer
func decodeType(jType string) int {
	// valid types
	types := []string{ "upload", "export", "teams" }

	// check whether command is present, and return an encoding
	for it, t := range types {
		if jType == t {
			return it
		}
	}

	return -1
}
