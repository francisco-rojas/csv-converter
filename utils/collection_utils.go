package utils

// Index returns the first index of the target string t, or -1 if no match is found.
func Index(c []string, t string) int {
	for i, v := range c {
		if v == t {
			return i
		}
	}
	return -1
}

// InArray checks if s is in the a array
func InArray(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
