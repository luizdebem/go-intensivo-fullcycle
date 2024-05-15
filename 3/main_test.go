package main

import "testing"

// rever benchmark no mac tbm pq no windows não rodou direito tbm

func BenchmarkGenerateLargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLargeString(1000)
	}
}
