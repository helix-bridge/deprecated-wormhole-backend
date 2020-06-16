package db

import (
	"fmt"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"sync"
)

type Supply struct {
	CirculatingSupply decimal.Decimal `json:"circulatingSupply"`
	TotalSupply       decimal.Decimal `json:"totalSupply"`
	MaxSupply         decimal.Decimal `json:"maxSupply"`
	Details           []*SupplyDetail `json:"details"`
}

type SupplyDetail struct {
	Network           string          `json:"network"`
	CirculatingSupply decimal.Decimal `json:"circulatingSupply"`
	TotalSupply       decimal.Decimal `json:"totalSupply"`
	MaxSupply         decimal.Decimal `json:"maxSupply"`
	Precision         int             `json:"precision"`
	Type              string          `json:"type,omitempty"`
	Contract          string          `json:"contract,omitempty"`
}

func CurrencySupply() *Supply {
	var supply Supply
	supply.MaxSupply = decimal.New(1, 12) // 10 billion

	ethSupply := ethSupply()
	tronSupply := tronSupply()

	supply.TotalSupply = ethSupply.TotalSupply.Add(tronSupply.TotalSupply)
	supply.CirculatingSupply = ethSupply.CirculatingSupply.Add(tronSupply.CirculatingSupply)
	supply.Details = []*SupplyDetail{ethSupply, tronSupply}

	return &supply
}

func ethSupply() *SupplyDetail {
	var supply SupplyDetail
	supply.Precision = 18
	precision := decimal.New(1, int32(supply.Precision))

	capDecimal := decimal.NewFromBigInt(parallel.RingEthSupply(), 0).Div(precision)
	supply.Network = "Ethereum"
	supply.CirculatingSupply = capDecimal.Sub(supply.filterBalance().Div(precision))
	supply.TotalSupply = capDecimal
	supply.MaxSupply = capDecimal
	supply.Type = "erc20"
	supply.Contract = "0x9469d013805bffb7d3debe5e7839237e535ec483"

	return &supply
}

func tronSupply() *SupplyDetail {
	var supply SupplyDetail
	supply.Precision = 18
	precision := decimal.New(1, int32(supply.Precision))

	capDecimal := decimal.NewFromBigInt(parallel.RingTronSupply(), 0).Div(precision)
	supply.Network = "Tron"
	supply.CirculatingSupply = capDecimal.Sub(supply.filterBalance().Div(precision))
	supply.TotalSupply = capDecimal
	supply.MaxSupply = capDecimal
	supply.Type = "trc20"
	supply.Contract = "TL175uyihLqQD656aFx3uhHYe1tyGkmXaW"

	return &supply
}

func (s *SupplyDetail) filterBalance() decimal.Decimal {
	var filterAddress = map[string][]string{
		"Tron": {"TDWzV6W1L1uRcJzgg2uKa992nAReuDojfQ",
			"TSu1fQKFkTv95U312R6E94RMdixsupBZDS",
			"TTW2Vpr9TCu6gxGZ1yjwqy7R79hEH8iscC",
		},
		"Ethereum": {"0x4710573b853fdd3561cb4f60ec9394f0155d5105",
			"0x7f23e4a473db3d11d11b43d90b34f8a778753e34",
			"0x649fdf6ee483a96e020b889571e93700fbd82d88",
		},
	}
	filter := filterAddress[s.Network]
	wg := sync.WaitGroup{}
	var sum decimal.Decimal
	for _, address := range filter {
		go func(address string) {
			defer wg.Done()
			switch s.Network {
			case "Tron":
				sum = sum.Add(decimal.NewFromBigInt(parallel.RingTronBalance(util.TrxBase58toHexAddress(address)), 0))
			case "Ethereum":
				sum = sum.Add(decimal.NewFromBigInt(parallel.RingEthBalance(address), 0))
				fmt.Println(sum, address)
			}
		}(address)
		wg.Add(1)
	}
	wg.Wait()
	return sum
}
