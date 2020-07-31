package main

// encode integer into a state string
func encodeState(command int) string {
	states := []string{ "start", "pause", "kill" }

	// if command value is illegal
	if command >= len(states) || command < 0 {
		return ""
	}

	return states[command]
}

// encode integer into a type string
func encodeType(jType int) string {
	// valid types
	types := []string{ "upload", "export", "teams" }

	// if command value is illegal
	if jType >= len(types) || jType < 0 {
		return ""
	}

	return types[jType]
}
