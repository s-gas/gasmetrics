package reading

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Reading struct {
	Value int       `json:"value"`
	Date  time.Time `json:"recorded_at"`
}

func Parse(msg string) (Reading, error) {
	if len(msg) == 0 {
		return Reading{}, errors.New("Parse: message is empty")
	}
	var r Reading
	var err error
	r.Value, err = strconv.Atoi(msg)
	if err != nil {
		return Reading{}, fmt.Errorf("Parse: %w", err)
	}
	r.Date = time.Now()
	return r, nil
}
