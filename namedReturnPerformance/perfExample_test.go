package main

import "testing"

func BenchmarkReturnTypeOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = declareReturnTypeOnly()
	}
}
func BenchmarkReturnNameTypeWithFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = declareReturnNameTypeWithFmt()
	}
}
func BenchmarkReturnNameTypeWithLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = declareReturnNameTypeWithLog()
	}
}
