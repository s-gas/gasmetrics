package handler

import (
	"fmt"
)

func validate(value int, lastEntry int, nextEntry int) error {
	if value <= lastEntry {
		return fmt.Errorf("validate: value must be greater than last entry (%v)", lastEntry)
	}
	if nextEntry != 0 && value > nextEntry {
		return fmt.Errorf("validate: value must be less than next entry (%v)", nextEntry)
	}
	return nil
}
