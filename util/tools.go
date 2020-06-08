package util

import (
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		return 0
	} else {
		return i
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
