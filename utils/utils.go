package utils

import (
	"strconv"
)

func CalculatePercent(success int, failure int) string {
	total := success + failure
	if total == 0 {
		return "0"
	}
	return strconv.Itoa((success * 100) / total)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
