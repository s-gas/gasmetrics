package handler

import (
	"fmt"
	"strconv"
)

func validate(value int, lastEntry int) (int, error) {
	if v <= lastEntry {
		return 0, fmt.Errorf("validate: value must be greater than last entry (%v)", lastEntry)
	}
	return v, nil
}
