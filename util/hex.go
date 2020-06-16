package util

import (
	"encoding/hex"
	"github.com/huandu/xstrings"
	"math/big"
	"strings"
)

func Padding(str string) string {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	return xstrings.RightJustify(str, 64, "0")
}

func BytesToHex(b []byte) string {
	c := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(c, b)
	return string(c)
}

func AddHex(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	if strings.HasPrefix(s, "0x") {
		return s
	}
	return strings.ToLower("0x" + s)
}

func AddTronPerfix(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	if strings.HasPrefix(s, "41") {
		return s
	}
	return strings.ToLower("41" + s)
}

func TrimHex(s string) string {
	return strings.TrimPrefix(s, "0x")
}

func TrimTronHex(s string) string {
	return strings.TrimPrefix(s, "41")
}

func U256(v string) *big.Int {
	v = strings.TrimPrefix(v, "0x")
	bn := new(big.Int)
	n, _ := bn.SetString(v, 16)
	return n
}

func LogAnalysis(log string) []string {
	log = strings.TrimPrefix(log, "0x")
	logLength := len(log)
	var logSlice []string
	for i := 0; i < logLength/64; i++ {
		logSlice = append(logSlice, log[i*64:(i+1)*64])
	}
	return logSlice
}
