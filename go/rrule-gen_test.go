package main

import "testing"

func BenchmarkGenerateEvents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateEvents()
	}
}
