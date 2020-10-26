package util

import (
	"encoding/json"
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToString(i interface{}) string {
	var val string
	switch i := i.(type) {
	case string:
		val = i
	case []byte:
		val = string(i)
	default:
		b, _ := json.Marshal(i)
		val = string(b)
	}
	return val
}
