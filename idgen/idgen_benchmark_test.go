package idgen

import (
	"testing"
)

// ✅ Benchmark: Generating IDs (default settings)
func BenchmarkGenerate(b *testing.B) {
	idGen := New(6, "", "")

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = idGen.Generate()
	}
}

// ✅ Benchmark: Generating IDs with long random parts (e.g., 16 characters)
func BenchmarkGenerate_LongRandomPart(b *testing.B) {
	idGen := New(16, "", "")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = idGen.Generate()
	}
}

// ✅ Benchmark: Generating IDs with a large number of concurrent goroutines
func BenchmarkGenerate_Concurrent(b *testing.B) {
	idGen := New(6, "", "")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = idGen.Generate()
		}
	})
}

// ✅ Benchmark: Base64 Encoding
func BenchmarkEncodeBase64(b *testing.B) {
	id := "240210120530XYZ9876"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeBase64(id)
	}
}

// ✅ Benchmark: Base58 Encoding
func BenchmarkEncodeBase58(b *testing.B) {
	id := "240210120530XYZ9876"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeBase58(id)
	}
}
