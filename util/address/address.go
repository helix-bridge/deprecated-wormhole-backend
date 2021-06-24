package address

import (
	"bytes"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/ss58"
	"regexp"
)

var (
	ethAddressRegex       = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	substrateAddressRegex = regexp.MustCompile(`^[0-9a-fA-F]{64}$`)
	numberRegexString     = regexp.MustCompile("^[0-9]+$")
)

const (
	GatewayReserved = "1"
)

var addressTag = map[string]string{
	GatewayReserved: "Reserved",
}

func Decode(address string, addressType int) string {
	if VerifyEthereumAddress(address) {
		return address
	}
	return ss58.Decode(address, addressType)
}

func Encode(accountID string, addressType int) string {
	if numberRegexString.MatchString(accountID) {
		return addressTag[accountID]
	}
	if VerifyEthereumAddress(accountID) {
		return util.AddHex(accountID)
	}
	return ss58.Encode(accountID, addressType)
}

func IsSubstrateSubAccount(accountID string) bool {
	accountID = util.TrimHex(accountID)
	bytesAccountID := util.HexToBytes(accountID)
	ZeroBytesSuffix := make([]byte, 8)
	if len(bytesAccountID) == 32 && bytes.HasSuffix(bytesAccountID, ZeroBytesSuffix) {
		return true
	}
	return false
}

func VerifySubstrateAddress(accountId string) bool {
	return substrateAddressRegex.MatchString(util.TrimHex(accountId))
}

func VerifyEthereumAddress(accountId string) bool {
	return ethAddressRegex.MatchString(util.AddHex(accountId))
}
