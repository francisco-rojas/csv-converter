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
