package slices

// Checks is slice contains int
func ContainsInt(slice []int, value int) bool {
	for i := range slice {
		if value == slice[i] {
			return true
		}
	}
	return false
}

// Checks if slice contains string
func ContainsString(slice []string, value string) bool {
	for i := range slice {
		if value == slice[i] {
			return true
		}
	}
	return false
}
