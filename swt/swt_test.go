package swt

import (
	"testing"
)

func BenchmarkSWTencode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeSWT(i)
	}
}

func BenchmarkSWTmulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := EncodeSWT(i)
		DecodeSWT(payload)
	}
}
