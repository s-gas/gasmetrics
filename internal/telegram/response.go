package telegram

func GetResponse(message string) string {
	if !isPositiveInteger(message) {
		return "Invalid input: enter a positive integer"
	} else {
		return "Input received"
	}
}
