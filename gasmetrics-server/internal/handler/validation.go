package handler

import (
	"fmt"
	"strconv"
)

func validate(message string, lastEntry int) (int, error) {
	v, err := strconv.Atoi(message)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}
	if v <= lastEntry {
		return 0, fmt.Errorf("validate: value must be greater than last entry (%v)", lastEntry)
	}
	return v, nil
}
