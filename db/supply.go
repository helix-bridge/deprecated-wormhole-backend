package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"strings"
	"sync"
)

type Supply struct {
	CirculatingSupply decimal.Decimal `json:"circulatingSupply"`
	TotalSupply       decimal.Decimal `json:"totalSupply"`
	BondLockBalance   decimal.Decimal `json:"bond_lock_balance"`
	TreasuryLockBalance decimal.Decimal `json:"treasury_lock_balance"`
	MaxSupply         decimal.Decimal `json:"maxSupply"`
	Details           []*SupplyDetail `json:"details"`
}

type SupplyDetail struct {
	Network           string          `json:"network"`
	CirculatingSupply decimal.Decimal `json:"circulatingSupply"`
	TotalSupply       decimal.Decimal `json:"totalSupply"`
	Precision         int             `json:"precision"`
	Type              string          `json:"type,omitempty"`
	Contract          string          `json:"contract,omitempty"`
}

type Currency struct {
	Code          string
	EthContract   string
	TronContract  string
	MaxSupply     decimal.Decimal
	FilterAddress map[string][]string
}

func RingSupply() *Supply {
	ring := Currency{
		Code:         "ring",
		EthContract:  config.Link.Ring,
		TronContract: config.Link.TronRing,
		MaxSupply:    decimal.New(1, 10),
	}
	ring.FilterAddress = map[string][]string{
		"Tron":     {"TDWzV6W1L1uRcJzgg2uKa992nAReuDojfQ", "TSu1fQKFkTv95U312R6E94RMdixsupBZDS", "TTW2Vpr9TCu6gxGZ1yjwqy7R79hEH8iscC"},
		"Ethereum": {"0x5FD8bCC6180eCd977813465bDd0A76A5a9F88B47", "0xfA4FE04f69F87859fCB31dF3B9469f4E6447921c", "0x7f23e4a473db3d11d11b43d90b34f8a778753e34", "0x7f23e4a473db3d11d11b43d90b34f8a778753e34"},
	}
	return ring.supply()
}

func KtonSupply() *Supply {
	kton := Currency{
		Code:         "kton",
		EthContract:  config.Link.Kton,
		TronContract: config.Link.TronKton,
	}
	return kton.supply()
}

// todo，need cache here
func (c *Currency) supply() *Supply {
	var supply Supply
	supply.MaxSupply = c.MaxSupply // 10 billion
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		ethSupply := c.ethSupply()
		supply.CirculatingSupply = supply.CirculatingSupply.Add(ethSupply.CirculatingSupply)
		supply.Details = append(supply.Details, ethSupply)
		wg.Done()
	}()
	go func() {
		tronSupply := c.tronSupply()
		supply.CirculatingSupply = supply.CirculatingSupply.Add(tronSupply.CirculatingSupply)
		supply.Details = append(supply.Details, tronSupply)
		wg.Done()
	}()
	go func() {
		supply.TreasuryLockBalance = c.TreasuryLock(100, 0, "councilMember")
		wg.Done()
	}()
	go func() {
		supply.TotalSupply, supply.BondLockBalance = c.TotalSupply()
		wg.Done()
	}()
	wg.Wait()

	if supply.MaxSupply.IsZero() {
		if c.Code == "kton"{
			supply.MaxSupply = supply.TotalSupply
		}else{
			for _, one := range supply.Details {
				supply.MaxSupply = supply.MaxSupply.Add(one.TotalSupply)
			}
		}
	}

	supply.CirculatingSupply = supply.TotalSupply.Sub(supply.BondLockBalance).Sub(supply.TreasuryLockBalance).
		Sub(supply.CirculatingSupply)
	return &supply
}

func (c *Currency) ethSupply() *SupplyDetail {
	var supply SupplyDetail
	supply.Precision = 18
	precision := decimal.New(1, int32(supply.Precision))

	capDecimal := decimal.NewFromBigInt(parallel.RingEthSupply(c.EthContract), 0).Div(precision)
	supply.Network = "Ethereum"
	supply.Contract = c.EthContract
	supply.CirculatingSupply = capDecimal.Sub(supply.filterBalance(c.FilterAddress).Div(precision))
	supply.TotalSupply = capDecimal
	supply.Type = "erc20"

	return &supply
}

func (c *Currency) tronSupply() *SupplyDetail {
	var supply SupplyDetail
	supply.Precision = 18
	precision := decimal.New(1, int32(supply.Precision))

	capDecimal := decimal.NewFromBigInt(parallel.RingTronSupply(c.TronContract), 0).Div(precision)
	supply.Contract = c.TronContract
	supply.Network = "Tron"
	supply.CirculatingSupply = capDecimal.Sub(supply.filterBalance(c.FilterAddress).Div(precision))
	supply.TotalSupply = capDecimal
	supply.Type = "trc20"

	return &supply
}

func (c *Currency) TreasuryLock(pageSize, pageIndex  int64, filter string) (decimal.Decimal)  {
	type AccountDetail struct {
		Balance decimal.Decimal `json:"balance"`
		BalanceLock decimal.Decimal `json:"balance_lock"`
		KtonBalance decimal.Decimal `json:"kton_balance"`
		KtonLock decimal.Decimal `json:"kton_lock"`
	}
	type AccountTokenRes struct {
		Data struct{
			Count int `json:"count"`
			List []AccountDetail `json:"list"`
		} `json:"data"`
	}

	params := make(map[string]interface{})
	params["row"] =  pageSize
	params["page"] = pageIndex
	params["filter"] = filter

	b, _ := json.Marshal(params)
	var res AccountTokenRes
	data, _ := util.PostWithJson(fmt.Sprintf("%s/api/scan/accounts", config.Link.SubscanHost), bytes.NewReader(b))
	util.UnmarshalAny(&res, data)

	var token decimal.Decimal

	for _, a := range res.Data.List{
		if c.Code == "ring"{
			token = token.Add(a.BalanceLock)
		}else{
			token = token.Add(a.BalanceLock)
		}
	}

	return token
}

func (c *Currency) TotalSupply() (decimal.Decimal, decimal.Decimal) {
	type TokenDetail struct {
		TotalIssuance decimal.Decimal `json:"total_issuance"`
		TokenDecimals int             `json:"token_decimals"`
		BondedLockedBalance decimal.Decimal `json:"bonded_locked_balance"`
	}
	type SubscanTokenRes struct {
		Data struct {
			Detail map[string]TokenDetail `json:"detail"`
		} `json:"data"`
	}
	var res SubscanTokenRes
	b, _ := util.HttpGet(fmt.Sprintf("%s/api/scan/token", config.Link.SubscanHost))
	util.UnmarshalAny(&res, b)
	detail := res.Data.Detail[strings.ToUpper(c.Code)]
	return detail.TotalIssuance.Div(decimal.New(1, int32(detail.TokenDecimals))),
		detail.BondedLockedBalance.Div(decimal.New(1, int32(detail.TokenDecimals)))
}

func (s *SupplyDetail) filterBalance(filterAddress map[string][]string) decimal.Decimal {
	filter := filterAddress[s.Network]
	wg := sync.WaitGroup{}
	var sum decimal.Decimal
	for _, address := range filter {
		go func(address string) {
			defer wg.Done()
			switch s.Network {
			case "Tron":
				sum = sum.Add(decimal.NewFromBigInt(parallel.RingTronBalance(s.Contract, util.TrxBase58toHexAddress(address)), 0))
			case "Ethereum":
				sum = sum.Add(decimal.NewFromBigInt(parallel.RingEthBalance(s.Contract, address), 0))
				fmt.Println(sum, address)
			}
		}(address)
		wg.Add(1)
	}
	wg.Wait()
	return sum
}
