package crypto

import "github.com/darwinia-network/link/pkg/go-ethereum/crypto/sha3"

// String string
func String(input interface{}) []byte {
	switch v := input.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	default:
		return []byte("")
	}
}

// Bool bool
func Bool(input interface{}) []byte {
	switch v := input.(type) {
	case bool:
		if v {
			return []byte{0x1}
		}
		return []byte{0x0}
	default:
		return []byte{0x0}
	}
}

// ConcatByteSlices concat byte slices
func ConcatByteSlices(arrays ...[]byte) []byte {
	var result []byte

	for _, b := range arrays {
		result = append(result, b...)
	}

	return result
}

// SoliditySHA3 solidity sha3
func SoliditySHA3(data ...[]byte) []byte {
	var result []byte

	hash := sha3.NewKeccak256()
	bs := ConcatByteSlices(data...)

	hash.Write(bs)
	result = hash.Sum(result)

	return result
}
