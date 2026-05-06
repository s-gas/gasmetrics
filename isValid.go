package main

import "strconv"

func isValid(message string) bool {
	if v, err := strconv.Atoi(message); err == nil {
		return v > 0
	} else {
		return false
	}
}
