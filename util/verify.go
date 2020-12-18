package util

import "regexp"

var (
	TronAddressRegex      = regexp.MustCompile(`^41[0-9a-fA-F]{40}$`)
	EthAddressRegex       = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	SubstrateAddressRegex = regexp.MustCompile(`[0-9a-fA-F]{64}$`)
)

func VerifyTronAddress(address string) bool {
	return TronAddressRegex.MatchString(address)
}

func VerifyEthAddress(address string) bool {
	return EthAddressRegex.MatchString(address)
}

func VerifySubstrateAddress(address string) bool {
	return SubstrateAddressRegex.MatchString(address)
}
