package benchmarks

import (
	"testing"

	"github.com/SmrutAI/pedantigo"
)

// ============================================================================
// Pedantigo Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Playground)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Validate_Simple validates an existing 5-field struct (bypass)
func Benchmark_Pedantigo_Validate_Simple(b *testing.B) {
	user := ValidUserPedantigo
	_ = pedantigo.Validate(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&user)
	}
}

// Benchmark_Pedantigo_Validate_Complex validates an existing nested struct (bypass)
func Benchmark_Pedantigo_Validate_Complex(b *testing.B) {
	order := ValidOrderPedantigo
	_ = pedantigo.Validate(&order) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&order)
	}
}

// Benchmark_Pedantigo_Validate_Large validates an existing 20+ field struct (bypass)
func Benchmark_Pedantigo_Validate_Large(b *testing.B) {
	config := ValidConfigPedantigo
	_ = pedantigo.Validate(&config) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&config)
	}
}

// ----------------------------------------------------------------------------
// Pedantigo Unique: UnmarshalMap (JSON → map → struct + validate)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_UnmarshalMap_Simple tests Pedantigo's unique JSON→map→struct flow
func Benchmark_Pedantigo_UnmarshalMap_Simple(b *testing.B) {
	_, _ = pedantigo.Unmarshal[UserPedantigo](ValidUserJSON) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Unmarshal[UserPedantigo](ValidUserJSON)
	}
}

// Benchmark_Pedantigo_UnmarshalMap_Complex tests Pedantigo's unique JSON→map→nested flow
func Benchmark_Pedantigo_UnmarshalMap_Complex(b *testing.B) {
	_, _ = pedantigo.Unmarshal[OrderPedantigo](ValidOrderJSON) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Unmarshal[OrderPedantigo](ValidOrderJSON)
	}
}

// ----------------------------------------------------------------------------
// Playground-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_UnmarshalDirect_Simple - Not applicable to Pedantigo
func Benchmark_Pedantigo_UnmarshalDirect_Simple(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern (json.Unmarshal + Struct)")
}

// Benchmark_Pedantigo_UnmarshalDirect_Complex - Not applicable to Pedantigo
func Benchmark_Pedantigo_UnmarshalDirect_Complex(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern (json.Unmarshal + Struct)")
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_New_Simple measures validator creation for simple struct
func Benchmark_Pedantigo_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.New[UserPedantigo]()
	}
}

// Benchmark_Pedantigo_New_Complex measures validator creation for nested struct
func Benchmark_Pedantigo_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.New[OrderPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Schema Generation (Pedantigo-only feature)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Schema_Uncached measures first-call schema generation
func Benchmark_Pedantigo_Schema_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := pedantigo.New[UserPedantigo]()
		_ = v.Schema()
	}
}

// Benchmark_Pedantigo_Schema_Cached measures cached schema retrieval
func Benchmark_Pedantigo_Schema_Cached(b *testing.B) {
	_ = pedantigo.Schema[UserPedantigo]() // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Schema[UserPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Marshal (Validate + JSON output)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Marshal_Simple measures validate + JSON marshal
func Benchmark_Pedantigo_Marshal_Simple(b *testing.B) {
	user := ValidUserPedantigo
	_, _ = pedantigo.Marshal(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Marshal(&user)
	}
}
