package observer

import (
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type EthTransaction struct {
	Last    int64                     `json:"last"`
	Address string                    `json:"address"`
	Method  []string                  `json:"method"`
	Result  *parallel.EtherscanResult `json:"result"`
}

func (e *EthTransaction) Do(o Observable) error {
	fmt.Println("find EthTransaction", e.Result)
	if e.Result == nil || !util.StringInSlice(e.Result.Topics[0], EthAvailableEvent) {
		return fmt.Errorf("empty transaction %s", e.Result)
	}
	return e.Redeem()
}

func (e *EthTransaction) Listen(o Observable) error {
	key := strings.Join(e.Method, ":")
	if e.Last == 0 {
		if b := util.GetCache(key); b != nil {
			e.Last = util.StringToInt64(string(b))
		} else {
			e.Last = 8028174
		}
	}
	go func() {
		for {
			if eventLog, _ := parallel.EtherscanLog(e.Last+1, e.Address, e.Method...); eventLog != nil {
				for _, result := range eventLog.Result {
					e.Last = util.U256(result.BlockNumber).Int64()
					e.Result = &result
					_ = o.notify(e)
				}
			}
			_ = util.SetCache(key, e.Last, 86400*7)
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

// https://github.com/darwinia-network/dj
func (e *EthTransaction) Redeem() error {
	logSlice := util.LogAnalysis(e.Result.Data)

	switch e.Result.Topics[0] {
	case BurnAndRedeem:
		currency := "ring"
		token := util.AddHex(e.Result.Topics[1][len(e.Result.Topics[1])-40:])
		from := util.AddHex(e.Result.Topics[2][len(e.Result.Topics[2])-40:])
		amount := decimal.NewFromBigInt(util.U256(logSlice[0]), 0)
		target := logSlice[3]
		if strings.EqualFold(token, config.Link.Kton) {
			currency = "kton"
		}
		return db.AddRedeemRecord(Eth, util.AddHex(e.Result.TransactionHash), from, target, currency, amount,
			int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()), "")

	case BurnAndRedeemDeposit:
		depositId := util.U256(e.Result.Topics[1]).Int64()
		from := util.AddHex(logSlice[0][len(logSlice[0])-40:])
		month := util.U256(logSlice[1]).Int64()
		startAt := util.U256(logSlice[2]).Int64()
		amount := decimal.NewFromBigInt(util.U256(logSlice[4]), 0)
		target := logSlice[7]
		deposit := map[string]int64{"deposit_id": depositId, "month": month, "start": startAt}

		return db.AddRedeemRecord(Eth, util.AddHex(e.Result.TransactionHash), from, target, "deposit", amount,
			int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()), util.ToString(deposit))

	case VerifyProof:
		blockNum := util.U256(logSlice[0]).Uint64()
		db.SetTokenTokenRegistrationConfirm(blockNum, util.AddHex(e.Result.TransactionHash))
		db.SetBackingLockConfirm(blockNum, util.AddHex(e.Result.TransactionHash))
		return nil

	case SetRootEvent:
		index := util.U256(logSlice[2]).Uint64()
		db.SetMMRIndexBestBlockNum(index)
	}
	return nil
}
