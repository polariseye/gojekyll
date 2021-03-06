package utils

// SearchStrings returns a bool indicating whether array contains the string.
// Unlike sort.SearchStrings, it does not require a sorted array.
// This is useful when the array length is low; else consider sorting it and using
// sort.SearchStrings, or creating a map[string]bool.
func SearchStrings(array []string, s string) bool {
	for _, item := range array {
		if item == s {
			return true
		}
	}
	return false
}

// StringSet creates a characteristic function map that tests for presence in an array.
func StringSet(array []string) map[string]bool {
	set := map[string]bool{}
	for _, s := range array {
		set[s] = true
	}
	return set
}
