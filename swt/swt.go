package swt

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
<<<<<<< HEAD
	"errors"
	"fmt"
=======
>>>>>>> parent of e9bcef7 (print error)
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Payload[key] = value
type ENCODE_KEY *string

var EXPIRE_TIME = time.Hour * 1
var config SWT_CONFIG = AutoConfig()

func NewEncodeKey() ENCODE_KEY {
	var sha256_key string
	var hash = sha256.New()

	today := time.Now()
	tommorow := today.AddDate(0, 0, rand.Intn(100))
	encode := (strconv.Itoa(today.Nanosecond()) + tommorow.String()) + today.String()

	hash.Write([]byte(encode))
	sha256_key = string(hash.Sum(nil))

	return &sha256_key
}

func AutoConfig() SWT_CONFIG {
	host, err := os.Hostname()
	if err != nil {
		host = "default-unsecure please change username due to security reason"
	}

	return SWT_CONFIG{
		EncodeKey: NewEncodeKey(),
		Signer:    host,
	}

}

// SWT_CONFIG is saved config on the servers that accept tokens
type SWT_CONFIG struct {
	EncodeKey ENCODE_KEY
	Signer    string
}

// Secure Web Token
// swt
type SWT struct {
	Typ     string
	Payload map[string]interface{}
	Expire  time.Time
}

func EncodeSWT(values map[string]interface{}) (Payload string, Err error) {
	now := time.Now()
	init := SWT{
		Typ:     "swt",
		Payload: values,
		Expire:  now.Add(EXPIRE_TIME),
	}
	// p = payload
	p := new(bytes.Buffer)

	enc := gob.NewEncoder(p)
<<<<<<< HEAD
	if err := enc.Encode(init); err != nil {
		fmt.Println(err)
		return "encode error", err
=======
	err := enc.Encode(init)
	if err != nil {
		return "encode error"
>>>>>>> parent of e9bcef7 (print error)
	}

	psec, err := encrypt(p.String())
	if err != nil {
<<<<<<< HEAD
		fmt.Println(err)
		return "encrypt error", err
=======
		return "encrypt error"
>>>>>>> parent of e9bcef7 (print error)
	}
	return psec, nil
}

func DecodeSWT(Payload string) (SWT, error) {
	var swt_cargo SWT
	punsec, err := decrypt(Payload)
	if err != nil {
		return swt_cargo, err
	}

	pbytes := bytes.NewBufferString(punsec)
	decoder := gob.NewDecoder(pbytes)
	if err = decoder.Decode(&swt_cargo); err != nil {
		return SWT{}, err
	}

	// if cargo expires then return nil
	if time.Now().After(swt_cargo.Expire) {
		return SWT{}, errors.New("Session time expired")
	}
	return swt_cargo, nil
}
