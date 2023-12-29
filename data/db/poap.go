package db

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// EthAddress is a common.Address wrapper
type EthAddress common.Address

// Scan scans the given value and assigns it to the EthAddress.
//
// The value parameter should be a string. It returns an error if the value
// cannot be assigned to the Address. The Address is assigned the value of the
// common.HexToAddress(str) after unmarshaling the JSONB value.
//
// It returns an error if the value cannot be assigned to the Address, otherwise
// it returns nil.
func (a *EthAddress) Scan(value any) error {
	if value == nil {
		return nil
	}
	
	str, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*a = EthAddress(common.BytesToAddress(str))
	return nil
}

// Value returns the value of the EthAddress as a driver.Value and an error.
//
// It converts the EthAddress to a common.Address and returns the hexadecimal representation as a string.
func (a *EthAddress) Value() (driver.Value, error) {
	return common.Address(*a).Hex(), nil
}

// EthAddrs is a slice of EthAddress
type EthAddrs []EthAddress

// Address returns a slice of common.Address containing the addresses from the EthAddrs receiver.
//
// No parameters.
// Return type: []common.Address.
func (e *EthAddrs) Address() []common.Address {
	var addrs []common.Address
	for _, v := range *e {
		addrs = append(addrs, common.Address(v))
	}
	return addrs
}

// BaseModel is the base model for all models.
type BaseModel struct {
	ID        uint64   `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:ID" json:"id"`
	CreatedAt JSONTime `gorm:"<-:false" json:"-"`
	UpdatedAt JSONTime `gorm:"<-:false" json:"-"`
}

// JSONTime for serialize time to json
type JSONTime struct {
	time.Time
}

// MarshalJSON implement json.Marshaler interface
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JSONTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value.In(time.UTC)}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// ToJSONTime convert time.Time to JSONTime
func ToJSONTime(t time.Time) JSONTime {
	return JSONTime{
		Time: t,
	}
}

// Accounts is the table of accounts.
type Accounts struct {
	BaseModel
	ProjectID     uint64 `gorm:"column:project_id;type:bigint(20) unsigned;default:0;comment:项目 ID;NOT NULL" json:"project_id"`
	ChainType     string `gorm:"column:chain_type;type:char(3);NOT NULL" json:"chain_type"`
	Name          string `gorm:"column:name;type:varchar(255);comment:链账户名称" json:"name"`
	NativeAddress string `gorm:"column:native_address;type:char(42);comment:地址;NOT NULL" json:"native_address"`
	HexAddress    string `gorm:"column:hex_address;type:char(46);comment:地址;NOT NULL" json:"hex_address"`
	PubKey        string `gorm:"column:pub_key;type:char(172);comment:公钥;NOT NULL" json:"pub_key"`
	PriKey        string `gorm:"column:pri_key;type:char(216);comment:私钥;NOT NULL" json:"pri_key"`
	AccIndex      uint64 `gorm:"column:acc_index;type:bigint(20) unsigned;default:0;comment:链地址偏移下标;NOT NULL" json:"acc_index"`
	Status        uint   `gorm:"column:status;type:tinyint(4) unsigned;default:0;comment:状态（1：已授权 2：未授权）;NOT NULL" json:"status"`
	TxID          uint64 `gorm:"column:tx_id;type:bigint(20) unsigned;default:0;comment:交易 ID;NOT NULL" json:"tx_id"`
	Algo          uint64 `gorm:"column:algo;type:bigint(20) unsigned;default:0;comment:加密算法, 1:secp256k1, 2:eth_secp256k1;NOT NULL" json:"algo"`
	OperationID   string `gorm:"column:operation_id;type:varchar(100);comment:操作 ID;NOT NULL" json:"operation_id"`
}

// TableName returns the table name of Accounts.
func (Accounts) TableName() string {
	return "t_accounts"
}

// AddressPageQuery retrieves a page of addresses from the Accounts table.
//
// It takes in two parameters:
//   - limit: an integer representing the maximum number of addresses to retrieve.
//   - offset: an integer representing the starting position of the addresses to retrieve.
//
// It returns two values:
//   - addrs: a slice of strings containing the retrieved addresses.
//   - err: an error, if any occurred during the retrieval process.
func (Accounts) AddressPageQuery(limit, offset int) (addrs EthAddrs, err error) {
	rs := db.Model(&Accounts{}).
		Select("hex_address").
		Limit(limit).
		Offset(offset).
		Order("id").
		Scan(&addrs)
	return addrs, rs.Error
}
