package swt

import (
	"testing"
)

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
