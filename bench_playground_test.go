package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
)

// playgroundValidator is the shared validator instance
var playgroundValidator = validator.New()

// ============================================================================
// Playground Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Pedantigo)
// ----------------------------------------------------------------------------

// Benchmark_Playground_Validate_Simple validates an existing 5-field struct
func Benchmark_Playground_Validate_Simple(b *testing.B) {
	user := ValidUserPlayground
	_ = playgroundValidator.Struct(user) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = playgroundValidator.Struct(user)
	}
}

// Benchmark_Playground_Validate_Complex validates an existing nested struct
func Benchmark_Playground_Validate_Complex(b *testing.B) {
	order := ValidOrderPlayground
	_ = playgroundValidator.Struct(order) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = playgroundValidator.Struct(order)
	}
}

// Benchmark_Playground_Validate_Large validates an existing 20+ field struct
func Benchmark_Playground_Validate_Large(b *testing.B) {
	config := ValidConfigPlayground
	_ = playgroundValidator.Struct(config) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = playgroundValidator.Struct(config)
	}
}

// ----------------------------------------------------------------------------
// Pedantigo-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Playground_UnmarshalMap_Simple - Not applicable to Playground
func Benchmark_Playground_UnmarshalMap_Simple(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature (JSON→map→struct)")
}

// Benchmark_Playground_UnmarshalMap_Complex - Not applicable to Playground
func Benchmark_Playground_UnmarshalMap_Complex(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature (JSON→map→struct)")
}

// ----------------------------------------------------------------------------
// Playground Unique: UnmarshalDirect (json.Unmarshal + Struct)
// ----------------------------------------------------------------------------

// Benchmark_Playground_UnmarshalDirect_Simple tests stdlib json.Unmarshal + Struct
func Benchmark_Playground_UnmarshalDirect_Simple(b *testing.B) {
	var user UserPlayground
	_ = json.Unmarshal(ValidUserJSON, &user)
	_ = playgroundValidator.Struct(user)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var u UserPlayground
		_ = json.Unmarshal(ValidUserJSON, &u)
		_ = playgroundValidator.Struct(u)
	}
}

// Benchmark_Playground_UnmarshalDirect_Complex tests stdlib json.Unmarshal + Struct for nested
func Benchmark_Playground_UnmarshalDirect_Complex(b *testing.B) {
	var order OrderPlayground
	_ = json.Unmarshal(ValidOrderJSON, &order)
	_ = playgroundValidator.Struct(order)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var o OrderPlayground
		_ = json.Unmarshal(ValidOrderJSON, &o)
		_ = playgroundValidator.Struct(o)
	}
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Playground_New_Simple measures validator creation
func Benchmark_Playground_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.New()
	}
}

// Benchmark_Playground_New_Complex - Playground uses same validator for all types
func Benchmark_Playground_New_Complex(b *testing.B) {
	b.Skip("Playground uses a single validator for all struct types")
}

// ----------------------------------------------------------------------------
// Schema Generation (Not supported)
// ----------------------------------------------------------------------------

// Benchmark_Playground_Schema_Uncached - Not supported by Playground
func Benchmark_Playground_Schema_Uncached(b *testing.B) {
	b.Skip("Playground does not support schema generation")
}

// Benchmark_Playground_Schema_Cached - Not supported by Playground
func Benchmark_Playground_Schema_Cached(b *testing.B) {
	b.Skip("Playground does not support schema generation")
}

// ----------------------------------------------------------------------------
// Marshal (Validate + JSON output)
// ----------------------------------------------------------------------------

// Benchmark_Playground_Marshal_Simple measures Struct + json.Marshal
func Benchmark_Playground_Marshal_Simple(b *testing.B) {
	user := ValidUserPlayground
	_ = playgroundValidator.Struct(user)
	_, _ = json.Marshal(user)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = playgroundValidator.Struct(user)
		_, _ = json.Marshal(user)
	}
}
