package handler

import (
	"fmt"
)

func validate(value int, lastEntry int) error {
	if value <= lastEntry {
		return fmt.Errorf("validate: value must be greater than last entry (%v)", lastEntry)
	}
	return nil
}
