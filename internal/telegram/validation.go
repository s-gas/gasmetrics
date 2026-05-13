package telegram

import "strconv"

func isPositiveInteger(message string) bool {
	if v, err := strconv.Atoi(message); err == nil {
		return v > 0
	} else {
		return false
	}
}
