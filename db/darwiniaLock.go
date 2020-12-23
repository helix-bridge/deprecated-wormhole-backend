package db

import (
	"encoding/binary"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type DarwiniaBackingLock struct {
	ExtrinsicIndex string          `json:"extrinsic_index" gorm:"primary_key;auto_increment:false"`
	CreatedAt      time.Time       `json:"-"`
	AccountId      string          `json:"account_id"`
	BlockNum       int             `json:"block_num"`
	BlockHash      string          `sql:"default: null;size:70" json:"block_hash" `
	RingValue      decimal.Decimal `json:"ring_value" sql:"type:decimal(30,0);"`
	KtonValue      decimal.Decimal `json:"kton_value" sql:"type:decimal(30,0);"`
	Target         string          `json:"target" sql:"default: null;size:100"`
	BlockTimestamp int             `json:"block_timestamp"`
	MMRIndex       uint            `json:"mmr_index"`
	MMRRoot        string          `json:"mmr_root"`
	Signatures     string          `json:"signatures" sql:"type:text;"`
	BlockHeader    string          `json:"block_header" sql:"type:text;"`
}

func CreateDarwiniaBacking(extrinsicIndex string, detail *parallel.ExtrinsicDetail) error {
	db := util.DB

	record := &DarwiniaBackingLock{
		ExtrinsicIndex: extrinsicIndex,
		BlockHash:      detail.BlockHash,
		BlockNum:       detail.BlockNum,
		BlockTimestamp: detail.BlockTimestamp,
		BlockHeader:    util.ToString(parallel.SubscanBlockHeader(uint(detail.BlockNum))),
	}
	for _, event := range detail.Event {
		switch event.EventId {
		case "LockRing":
			record.AccountId = util.AddHex(util.ToString(event.Params[0].Value))
			record.Target = util.ToString(event.Params[1].Value)
			record.RingValue = util.DecimalFromInterface(event.Params[3].Value)
		case "LockKton":
			record.AccountId = util.AddHex(util.ToString(event.Params[0].Value))
			record.Target = util.ToString(event.Params[1].Value)
			record.KtonValue = util.DecimalFromInterface(event.Params[3].Value)
		case "NewMMRRoot":
			record.MMRIndex = uint(util.StringToInt(util.ToString(event.Params[0].Value)))
		}
	}

	query := db.Create(&record)
	return query.Error
}

type MerkleMountainRangeRootLog struct {
	ParentMmrRoot string `json:"parent_mmr_root"`
}

func DarwiniaBackingLocks(accountId string, page, row int) ([]DarwiniaBackingLock, int) {
	var list []DarwiniaBackingLock
	var count int
	util.DB.Model(DarwiniaBackingLock{}).Where("account_id = ?", accountId).Count(&count)
	util.DB.Where("account_id = ?", accountId).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)

	for index, lock := range list {
		if lock.MMRRoot != "" {
			continue
		}
		logs := parallel.SubscanLogs(lock.MMRIndex + 1)
		for _, logData := range logs {
			if strings.ToLower(logData.LogType) == "other" {

				var merkleMountainRangeRootLog *MerkleMountainRangeRootLog
				util.UnmarshalAny(&merkleMountainRangeRootLog, logData.Data)
				if merkleMountainRangeRootLog != nil {
					_ = lock.setMMRRoot(merkleMountainRangeRootLog.ParentMmrRoot)
					list[index].MMRRoot = merkleMountainRangeRootLog.ParentMmrRoot
				}
			}
		}
	}
	return list, count
}

func (l *DarwiniaBackingLock) setMMRRoot(mmrRoot string) error {
	query := util.DB.Model(l).Update(DarwiniaBackingLock{MMRRoot: mmrRoot})
	return query.Error
}

type MMRSignedSignature struct {
	Col2 string `json:"col2"`
}

func MMRRootSigned(eventParams []parallel.EventParam) error {
	mmrIndex := util.IntFromInterface(eventParams[0].Value)
	mmrRoot := util.ToString(eventParams[1].Value)

	var signatures []MMRSignedSignature
	util.UnmarshalAny(&signatures, eventParams[2].Value)

	var signatureList []string
	for _, signature := range signatures {
		signatureList = append(signatureList, signature.Col2)
	}

	query := util.DB.Model(DarwiniaBackingLock{}).Where("mmr_index = ?", mmrIndex).
		Update(DarwiniaBackingLock{MMRRoot: util.AddHex(mmrRoot), Signatures: strings.Join(signatureList, ",")})

	return query.Error
}

func BuilderHeader(header *parallel.BlockHeader) string {
	b := strings.Builder{}
	b.WriteString("0x")
	b.WriteString(util.TrimHex(header.ParentHash))
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(header.BlockNumber))
	b.WriteString(util.BytesToHex(bs))
	b.WriteString(util.TrimHex(header.StateRoot))
	b.WriteString(util.TrimHex(header.ExtrinsicsRoot))
	bs = make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(len(header.Digest)<<2))
	b.WriteString(util.BytesToHex(bs[0:1]))
	for _, log := range header.Digest {
		b.WriteString(util.TrimHex(log))
	}
	return b.String()
}
