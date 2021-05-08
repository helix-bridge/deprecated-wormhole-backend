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

	RingBuildInEvent     = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("RingBuildInEvent(address,address,uint256,bytes)"))))
	KtonBuildInEvent     = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("KtonBuildInEvent(address,address,uint256,bytes)"))))
	BurnAndRedeem        = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("BurnAndRedeem(address,address,uint256,bytes)"))))
	BurnAndRedeemDeposit = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("BurnAndRedeem(uint256,address,uint48,uint48,uint64,uint128,bytes)"))))
	VerifyProof          = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("VerifyProof(uint32)"))))
	SetRootEvent         = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("SetRootEvent(address,bytes32,uint256)"))))
	BackingLock         = util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("BackingLock(address,address,address,uint256,address,uint256)"))))

	EthAvailableEvent = []string{
		RingBuildInEvent,
		KtonBuildInEvent,
		BurnAndRedeem,
		BurnAndRedeemDeposit,
		VerifyProof,
		SetRootEvent,
		BackingLock,
	}
)
