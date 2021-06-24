package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func HexToBytes(s string) []byte {
	s = strings.TrimPrefix(s, "0x")
	c := make([]byte, hex.DecodedLen(len(s)))
	_, _ = hex.Decode(c, []byte(s))
	return c
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

// Convert int64, uint64, float64, string to int, return 0 if other types
func IntFromInterface(i interface{}) int {
	switch i := i.(type) {
	case int:
		return i
	case int64:
		return int(i)
	case uint64:
		return int(i)
	case float64:
		return int(i)
	case string:
		return StringToInt(i)
	}
	return 0
}

func Debug(i interface{}) {
	var val string
	switch i := i.(type) {
	case string:
		val = i
	case []byte:
		val = string(i)
	case error:
		val = i.Error()
	default:
		b, _ := json.MarshalIndent(i, "", "  ")
		val = string(b)
	}
	fmt.Println(val)
}

func DecimalFromInterface(i interface{}) decimal.Decimal {
	switch i := i.(type) {
	case int:
		return decimal.New(int64(i), 0)
	case int64:
		return decimal.New(i, 0)
	case uint64:
		return decimal.New(int64(i), 0)
	case float64:
		return decimal.NewFromFloat(i)
	case string:
		r, _ := decimal.NewFromString(i)
		return r
	}
	return decimal.Zero
}
