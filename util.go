package panel

// util.go is meant for all miscellaneous contructs
// that are not exported and exist only to work
// behind the scenes

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
