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

func UnmarshalAny(r interface{}, raw interface{}) {
	switch raw := raw.(type) {
	case string:
		_ = json.Unmarshal([]byte(raw), &r)
	case []uint8:
		_ = json.Unmarshal(raw, &r)
	default:
		b, _ := json.Marshal(raw)
		_ = json.Unmarshal(b, r)
	}
}

// Convert int64, uint64, float64, string to int, return 0 if other types
func UInt64FromInterface(i interface{}) uint64 {
	switch i := i.(type) {
	case int:
		return uint64(i)
	case int64:
		return uint64(i)
	case uint64:
		return i
	case float64:
		return uint64(i)
	case string:
		return uint64(StringToInt(i))
	}
	return 0
}
