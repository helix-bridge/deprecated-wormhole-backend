package util

import (
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
