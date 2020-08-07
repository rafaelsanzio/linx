package utils

// IsInArray verify if some string has already in array
func IsInArray(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
