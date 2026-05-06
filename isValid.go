package main

import "strconv"

func isValid(message string) bool {
	if _, err := strconv.Atoi(message); err == nil {
		return true
	} else {
		return false
	}
}
