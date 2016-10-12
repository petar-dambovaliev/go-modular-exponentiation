package main

import "testing"

func BenchmarkModExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModExp(123, 456, 89)
	}
}

func BenchmarkModExpWithSquaring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModExpWithSquaring(123, 456, 89)
	}
}

func BenchmarkModExpGoBigInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModExpGoBigInteger(123, 456, 89)
	}
}

func BenchmarkModExpGoBigIntegerExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModExpGoBigIntegerExp(123, 456, 89)
	}
}
