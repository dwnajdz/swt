package swt

import (
	"bytes"
	"math"
	"testing"
)

func TestFunctionality(t *testing.T) {
	type TestStruct struct {
		MyBool       bool
		IntElement   int
		Binary       uint64
		FloatElement float64
		StrElement   string
		IntArr       []int
		StrArr       []string
		MyByte       []byte
		SomeMap      map[string]interface{}
	}

	myStruct := TestStruct{
		MyBool:       true,
		IntElement:   math.MaxInt,
		Binary:       math.MaxUint64,
		FloatElement: math.MaxFloat64,
		StrElement:   "sOmE TeSt StRiNG",
		IntArr:       []int{math.MaxInt, math.MinInt},
		StrArr:       []string{"001010212", "\t", "\n"},
		MyByte:       []byte("BYTES STRING \n"),
		SomeMap:      make(map[string]interface{}),
	}

	t.Log("Test struct: ", myStruct)
	encoded := EncodeSWTcustom(myStruct)
	decoded := DecodeSWT(encoded)
	decoded_struct := decoded.Payload.(TestStruct)

	t.Log("Encoded message: ", encoded)
	t.Log("Decoded: ", decoded_struct)
}

func TestGoTypes(t *testing.T) {
	myBts := bytes.NewBuffer([]byte("UNCODE THIS TEST 1")).Bytes()
	encoded := EncodeSWTcustom(myBts)

	decoded := DecodeSWT(encoded)
	decoded_byte := decoded.Payload.([]byte)

	decodedStr := string(decoded_byte)
	expected := "UNCODE THIS TEST 1"

	t.Log("Encoded message: ", encoded)
	t.Log("Decoded: ", decoded_byte)
	if decodedStr == expected {
		t.Logf("PASSED! \n, Expected: %s; Actual value: %s", expected, decoded_byte)
	} else {
		t.Fatalf("Expected: %s; Actual value: %s", expected, decoded_byte)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeSWT(i)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := EncodeSWT(i)
		DecodeSWT(payload)
	}
}
