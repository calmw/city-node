package utils

import (
	"crypto/sha256"
	"errors"
	"reflect"
	"strconv"
)

var (
	ErrDecodeLength = errors.New("base58 decode length error")
	ErrDecodeCheck  = errors.New("base58 check failed")
)

// StringToInt64 string to int64
func StringToInt64(s string) int64 {
	int64Num, _ := strconv.ParseInt(s, 10, 64)
	return int64Num
}

// StringToInt string to int
func StringToInt(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}

// Int64ToString int64 to string
func Int64ToString(n int64) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}

// Int32ToString int32 to string
func Int32ToString(n int32) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}

// StringToInt32 int32 to string
func StringToInt32(s string) int32 {
	var j int32
	int10, _ := strconv.ParseInt(s, 10, 32)
	j = int32(int10)
	return j
}

// Int64ToInt int64 to int
func Int64ToInt(n int64) int {
	strInt64 := strconv.FormatInt(n, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

func Byte32ToByteSlice(b32 [32]byte) []byte {
	var bytes []byte
	for i := 0; i < 32; i++ {
		bytes = append(bytes, b32[i])
	}
	return bytes
}

func ByteSliceToByte32(byteSlice []byte) [32]byte {
	var bytes32 [32]byte
	for i := 0; i < len(byteSlice); i++ {
		bytes32[i] = byteSlice[i]
	}
	return bytes32
}

// Hash implements crypto sha256 hash algorithm.
func Hash(s []byte) []byte {
	h := sha256.New()
	// Hash.Write never returns an error per godoc
	h.Write(s)
	return h.Sum(nil)
}

func StructToStringMap(obj interface{}) map[string]string {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).String()
	}
	return data
}
