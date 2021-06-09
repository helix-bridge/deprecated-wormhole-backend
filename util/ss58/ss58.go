package ss58

import (
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/base58"
	"golang.org/x/crypto/blake2b"
)

func Decode(address string, addressType int) string {
	checksumPrefix := []byte("SS58PRE")
	addressDecode := base58.Decode(address)
	if len(addressDecode) == 0 {
		return ""
	}

	ss58Format := addressDecode[0]
	if ss58Format != byte(addressType) {
		return ""
	}
	var checksumLength int
	if util.IntInSlice(len(addressDecode), []int{3, 4, 6, 10}) {
		checksumLength = 1
	} else if util.IntInSlice(len(addressDecode), []int{5, 7, 11, 34, 35}) {
		checksumLength = 2
	} else if util.IntInSlice(len(addressDecode), []int{8, 12}) {
		checksumLength = 3
	} else if util.IntInSlice(len(addressDecode), []int{9, 13}) {
		checksumLength = 4
	} else if util.IntInSlice(len(addressDecode), []int{14}) {
		checksumLength = 5
	} else if util.IntInSlice(len(addressDecode), []int{15}) {
		checksumLength = 6
	} else if util.IntInSlice(len(addressDecode), []int{16}) {
		checksumLength = 7
	} else if util.IntInSlice(len(addressDecode), []int{17}) {
		checksumLength = 8
	} else {
		return ""
	}
	bss := addressDecode[0 : len(addressDecode)-checksumLength]
	checksum, _ := blake2b.New(64, []byte{})
	w := append(checksumPrefix[:], bss[:]...)
	_, err := checksum.Write(w)
	if err != nil {
		return ""
	}

	h := checksum.Sum(nil)
	if util.BytesToHex(h[0:checksumLength]) != util.BytesToHex(addressDecode[len(addressDecode)-checksumLength:]) {
		return ""
	}
	return util.BytesToHex(addressDecode[1 : len(addressDecode)-checksumLength])
}

func Encode(address string, addressType int) string {
	checksumPrefix := []byte("SS58PRE")
	addressBytes := util.HexToBytes(address)
	if addressType < 0 || addressType > 16383 {
		return ""
	}
	var checksumLength int
	if util.IntInSlice(len(addressBytes), []int{32, 33}) {
		checksumLength = 2
	} else if util.IntInSlice(len(addressBytes), []int{1, 2, 4, 8}) {
		checksumLength = 1
	} else {
		return ""
	}
	addressFormatPrefix := []byte{byte(addressType)}[:]
	addressFormat := append(addressFormatPrefix, addressBytes[:]...)
	checksum, _ := blake2b.New(64, []byte{})
	w := append(checksumPrefix[:], addressFormat[:]...)
	_, err := checksum.Write(w)
	if err != nil {
		return ""
	}

	h := checksum.Sum(nil)
	b := append(addressFormat[:], h[:checksumLength][:]...)
	return base58.Encode(b)
}
