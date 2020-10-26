package observer

import (
	"encoding/hex"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
)

const (
	ring = "ring"
	kton = "kton"
	Eth  = "eth"
	Tron = "tron"
)

var (
	TronAvailableEvent = []string{
		"RingBuildInEvent",
		"KtonBuildInEvent",
	}

	EthAvailableEvent = []string{
		util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("RingBuildInEvent(address,address,uint256,bytes)")))),
		util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("KtonBuildInEvent(address,address,uint256,bytes)")))),
		util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("BurnAndRedeem(address,address,uint256,bytes)")))),
		util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("BurnAndRedeem(uint256,address,uint48,uint48,uint64,uint128,bytes)")))),
	}
)
