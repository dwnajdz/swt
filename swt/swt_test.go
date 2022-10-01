package swt_test

import (
	"fmt"
	"swt/swt"
	"testing"
)

/*
	m := make(map[string]interface{})
	for i := 0; i < 50; i++ {
		m[strconv.Itoa(i)] = i
	}
	//fmt.Println(swt.EncodeSWT(m))

	payload := swt.EncodeSWT(m)
	decoded := swt.DecodeSWT(payload)
	fmt.Println(decoded.Payload["1"])
*/

func BenchmarkSWTcreation(b *testing.B) {
	m := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		m["i"] = i
	}
	fmt.Println(swt.EncodeSWT(m))
}
