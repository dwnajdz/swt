package swt

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Payload[key] = value
type ENCODE_KEY *[]byte

// SWT_CONFIG is saved config on the servers that accept tokens
type SWT_CONFIG struct {
	EncodeKey ENCODE_KEY
	Signer    string
}

// Secure Web Token
// swt
type SWT struct {
	Payload interface{}
	Expire  time.Time
}

var EXPIRE_TIME = time.Hour * 1
var config SWT_CONFIG = AutoConfig()

func NewEncodeKey() ENCODE_KEY {
	var sha256_key []byte
	var hash = sha256.New()

	today := time.Now()
	tommorow := today.AddDate(0, 0, rand.Intn(100))
	encode := (strconv.Itoa(today.Nanosecond()) + tommorow.String()) + today.String()

	hash.Write([]byte(encode))
	sha256_key = hash.Sum(nil)

	return &sha256_key
}

func AutoConfig() SWT_CONFIG {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		host = "default-unsecure please change username due to security reason"
	}

	return SWT_CONFIG{
		EncodeKey: NewEncodeKey(),
		Signer:    host,
	}
}

// While you are using custom types like struct or etc
// use EncodeSWTcustom to register this type
//
// In EncodeSWT you can use default types like:
// - int, float, string, []string, []list, bool etc.
func EncodeSWT(value interface{}) (Payload string) {
	now := time.Now()
	init := SWT{
		Payload: value,
		Expire:  now.Add(EXPIRE_TIME),
	}
	// p = payload
	p := new(bytes.Buffer)

	enc := gob.NewEncoder(p)
	err := enc.Encode(init)
	if err != nil {
		fmt.Println(err)
		return "encode error"
	}

	psec, err := encrypt(p.Bytes(), *config.EncodeKey)
	if err != nil {
		fmt.Println(err)
		return "encrypt error"
	}
	return psec
}

// In EncodeSWT you can use custom types like:
// struct, map, interface or any custom named type.
//
// CANNOT USE NON EXPORTED TYPES FOR EXAMPLE: bytes.Buffer, time.
func EncodeSWTcustom(value interface{}) (Payload string) {
	now := time.Now()
	init := SWT{
		Payload: value,
		Expire:  now.Add(EXPIRE_TIME),
	}
	// p = payload
	p := new(bytes.Buffer)

	// if it is custom value something like struct or etc
	// register this value with gob
	gob.Register(value)
	enc := gob.NewEncoder(p)
	err := enc.Encode(init)
	if err != nil {
		fmt.Println(err)
		return "encode error"
	}

	psec, err := encrypt(p.Bytes(), *config.EncodeKey)
	if err != nil {
		fmt.Println(err)
		return "encrypt error"
	}
	return psec
}

func DecodeSWT(Payload string) SWT {
	var swt_cargo SWT
	punsec, err := decrypt(Payload, *config.EncodeKey)
	if err != nil {
		fmt.Println(err)
		return swt_cargo
	}

	pbytes := bytes.NewBufferString(punsec)
	decoder := gob.NewDecoder(pbytes)
	if err = decoder.Decode(&swt_cargo); err != nil {
		fmt.Println(err)
		return SWT{}
	}

	// if cargo expires then return nil
	if time.Now().After(swt_cargo.Expire) {
		return SWT{}
	}
	return swt_cargo
}

func (token SWT) IsPayloadNil() bool {
	return token.Payload == nil
}
