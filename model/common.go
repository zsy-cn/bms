package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringArray 字符串数组
type StringArray []string

// Uint64Array 整型数组
type Uint64Array []uint64

// Float64Array 整型数组
type Float64Array []float64

// Scan 解析前操作
func (sa *StringArray) Scan(data interface{}) (err error) {
	return scan(data, sa)
}
// Value 存储前操作
func (sa *StringArray) Value() (driver.Value, error) {
	return value(sa)
}

// Scan 解析前操作
func (ia *Uint64Array) Scan(data interface{}) (err error) {
	return scan(data, ia)
}
// Value 存储前操作
func (ia *Uint64Array) Value() (driver.Value, error) {
	result, err := json.Marshal(ia)
	return string(result), err
}

// Scan 解析前操作
func (fa *Float64Array) Scan(data interface{}) (err error) {
	return scan(data, fa)
}
// Value 存储前操作
func (fa *Float64Array) Value() (driver.Value, error) {
	result, err := json.Marshal(fa)
	return string(result), err
}

/********************************************************/
func scan(data interface{}, obj interface{}) (err error) {
	if data == nil {
		return
	}

	var byteData []byte
	switch values := data.(type) {
	case []byte:
		byteData = values
	case string:
		byteData = []byte(values)
	default:
		err = errors.New("unsupported driver")
		return
	}
	err = json.Unmarshal(byteData, obj)
	return
}

func value(obj interface{}) (driver.Value, error) {
	result, err := json.Marshal(obj)
	return string(result), err
}
