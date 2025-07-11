package functions

/// IsValidInput checks if the input string contains only printable ASCII characters (from 32 to 126).
func IsValidInput(input string) bool {
	for _, val := range input {
		if val < 32 || val > 126 {
			return false
		}
	}
	return true
}