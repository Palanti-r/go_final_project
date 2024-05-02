package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	originalDate, err := time.Parse("20060102", date)
	if err != nil {

		return "", err
	}
	if repeat == "" {
		return "", errors.New("repeat rule is empty")
	}
	repeatRules := strings.Split(repeat, " ")

	switch repeatRules[0] {
	case "d":
		if len(repeatRules) != 2 {
			return "", errors.New("invalid repeat rule format")
		}
		d, err := strconv.Atoi(repeatRules[1])
		if err != nil || d > 400 || d < 0 {
			return "", errors.New("wrong number of days")
		}
		for {
			originalDate = originalDate.AddDate(0, 0, d)
			if !originalDate.Equal(now) && !originalDate.Before(now) {
				break
			}
		}
	case "y":
		for {
			originalDate = originalDate.AddDate(1, 0, 0)
			if !originalDate.Equal(now) && !originalDate.Before(now) {
				break
			}
		}
	default:
		return "", errors.New("wrong repeat rule format")
	}
	return originalDate.Format("20060102"), nil
}
