package functions


func IsValidInput(input string) bool {
	for _, val := range input {
		if val < 32 || val > 126 {
			return false
		}
	}
	return true
}