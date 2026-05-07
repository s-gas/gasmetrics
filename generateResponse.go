package main

func generateResponse(message string) string {
	if !isValid(message) {
		return "Invalid input: enter a positive integer"
	} else {
		return "Input received"
	}
}
