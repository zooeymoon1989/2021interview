package gopherCon

import "testing"

type T struct {
	X int32
}

var global interface{}

func BenchmarkAllocOnHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		global = &T{}
	}
}

func BenchmarkAllocOnStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		local := T{}
		_ = local
	}
}
